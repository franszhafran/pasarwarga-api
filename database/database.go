package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	dsn := "root:654321@tcp(127.0.0.1:3309)/pasarwarga_article?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
}

func GetConnection() *gorm.DB {
	if db == nil {
		InitDB()
	}
	return db
}
