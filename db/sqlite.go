package db

import (
	// "database/sql"
	// "log"

	// _ "github.com/mattn/go-sqlite3"
	"database/sql"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func InitSqlite(path string) *sql.DB {
	db, err := gorm.Open("sqlite3", path)
	if err != nil {
		panic("连接数据库失败")
	}

	return db
}
