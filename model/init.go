package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"minepin-backend/config"
)

type Database struct {
	DB *gorm.DB
}

var DB *Database

func openSqliteDB(path string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&UserModel{})
	return db
}

func (db *Database) Init() {
	DB = &Database{
		DB: openSqliteDB(config.GetString("db.addr")),
	}
}

