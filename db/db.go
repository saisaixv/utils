package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var Client *DB

type DB struct {
	Client *gorm.DB
}

func Release() error {
	if Client == nil {
		return fmt.Errorf("Db CLient is nil")
	}

	if err := Client.Client.Close(); err != nil {
		return err
	}

	return nil
}
