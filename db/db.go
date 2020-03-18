package db

import (
	"database/sql"
	"fmt"
)

var Client *DB

type DB struct {
	Client *sql.DB
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
