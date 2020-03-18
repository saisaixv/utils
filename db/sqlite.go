package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func InitSqlite(path string) *gorm.DB {
	db, err := gorm.Open("sqlite3", path)
	if err != nil {
		panic("连接数据库失败")
	}

	return db
}
