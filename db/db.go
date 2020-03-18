package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func Release() error {
	if DB == nil {
		return fmt.Errorf("Db CLient is nil")
	}

	if err := DB.Close(); err != nil {
		return err
	}
	DB = nil
	return nil
}
