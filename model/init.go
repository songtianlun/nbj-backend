package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"minepin-backend/config"
	"minepin-backend/pkg/logger"
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
	err = db.AutoMigrate(&UserModel{})
	if err != nil {
		panic("fail to auto migrate db: " + err.Error())
	}
	return db
}

func openMySqlDB(path string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(path), &gorm.Config{})
	if err != nil {
		panic("failed to connect with MySQL - " + err.Error())
	}
	if err = db.AutoMigrate(&UserModel{}); err != nil {
		panic("fail to auto migrate [User] db: " + err.Error())
	}
	//if err = db.AutoMigrate(&UserATokenModel{}); err != nil {
	//	panic("fail to auto migrate [UserAccessToken] db: " + err.Error())
	//}
	if err = db.AutoMigrate(&UserRTokenModel{}); err != nil {
		panic("fail to auto migrate [UserRTokenModel] db: " + err.Error())
	}
	if err = db.AutoMigrate(&UserLoginLog{}); err != nil {
		panic("fail to auto migrate [UserLoginLog] db: " + err.Error())
	}
	if err = db.AutoMigrate(&UserPrefModel{}); err != nil {
		panic("fail to auto migrate [UserPreFSModel] db: " + err.Error())
	}
	return db
}

func (db *Database) Init() {
	var gdb *gorm.DB
	switch config.GetMinePinDbType() {
	case "sqlite3":
		gdb = openSqliteDB(config.GetMinePinDbAddr())
	case "mysql":
		gdb = openMySqlDB(
			fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				config.GetMinePinDbUserName(),
				config.GetMinePinDbPassWord(),
				config.GetMinePinDbAddr(),
				config.GetMinePinDbName()))
	}
	logger.InfoF("connected to %s with %s",
		config.GetMinePinDbAddr(), config.GetMinePinDbType())
	DB = &Database{
		DB: gdb,
	}
}
