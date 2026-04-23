package common

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() *gorm.DB {
	dbPath := "./data/gorm.db"

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		fmt.Println("db err: (Init) ", err)
	}

	sqlDB, err := db.DB()

	if err != nil {
		fmt.Println("db err: (Init - get sql.DB)", err)
	} else {
		sqlDB.SetMaxIdleConns(10)
	}

	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
