package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func AesEncrypt(value, key string) string {

	origData := []byte(value)
	k := []byte(key)

	block, _ := aes.NewCipher(k)
	blockSize := block.BlockSize()

	origData = PKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return base64.StdEncoding.EncodeToString(crypted)
}

func AesDecrypt(crypted string, key string) string {
	cryptedByte, _ := base64.StdEncoding.DecodeString(crypted)
	k := []byte(key)
	block, _ := aes.NewCipher(k)
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, k[:blocksize])
	orig := make([]byte, len(cryptedByte))

	blockMode.CryptBlocks(orig, cryptedByte)

	orig = PKCS7UnPadding(orig)
	return string(orig)
}

func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
