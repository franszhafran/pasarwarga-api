package migration

import (
	"fmt"
	"pasarwarga-service-api/src/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Article struct {
	ID         uint
	Title      string
	CategoryID int
	Content    string
	gorm.Model
}

func Migrate() {
	fmt.Println("Running migrations...")
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:654321@tcp(127.0.0.1:3309)/pasarwarga_article?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(models.Article{})
}
