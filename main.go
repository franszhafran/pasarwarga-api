package main

import (
	"os"

	"pasarwarga-service-api/migration"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB(db *gorm.DB) {
	dsn := "root:654321@tcp(127.0.0.1:3309)/pasarwarga_article?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
}

func GetConnection() *gorm.DB {
	return db
}

func main() {
	command := os.Args[1]
	initDB(db)

	if command == "migrate" {
		migration.Migrate()
	} else {
		r := gin.Default()

		initRoutes(r)

		r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	}
}
