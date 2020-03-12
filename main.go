package main

import (
	"encoding/json"
	"fmt"

	"github.com/saisaixv/utils/crypto"
)

type testTask struct{}

func (t *testTask) Run() {
	fmt.Println("hello world")
}

func main() {
	// crontab:=cron.NewCrontab()
	// task:=&testTask{}
	// if err:=crontab.AddByJob("1","* * * * *",task);err!=nil{
	// 	fmt.Printf("error to add crontab task:%s",err)
	// 	os.Exit(-1)
	// }

	// taskFunc:=func(){
	// 	fmt.Println("hello world")
	// }

	// if err:=crontab.AddByFunc("2","* * * * *",taskFunc);err!=nil{
	// 	fmt.Printf("error to add crontab task:%s",err)
	// 	os.Exit(-1)
	// }
	// crontab.Start()
	// select{}

	// redis.InitPool("192.168.10.168", 6379, "caton", 1, 30, 30, 30, 30)

	// err := redis.DoStrSet("a", "AAAA")
	// if err != nil {
	// 	fmt.Printf("set err: %s\n", err.Error())
	// }

	// ret, err := redis.DoStrGet("a")
	// if err != nil {
	// 	fmt.Printf("get err: %s\n", err.Error())
	// }
	// fmt.Println(ret)

	proxy := &Proxy{
		Host:     "192.168.10.168",
		Port:     8388,
		Password: "caton",
		Method:   "aes-256-cfb",
	}

	bArr, err := json.Marshal(proxy)
	if err != nil {
		fmt.Printf("Login:json.Marshal error = %s\n", err.Error())
	}

	ret := crypto.AesEncrypt(string(bArr), "abcdefghabcdefgh")
	fmt.Println(ret)

	ret = crypto.AesDecrypt(ret, "abcdefghabcdefgh")
	fmt.Println(ret)
}

type Proxy struct {
	Host     string
	Port     int
	Password string
	Method   string
}
