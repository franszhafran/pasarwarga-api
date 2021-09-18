package migration

import (
	"fmt"
	"pasarwarga-service-api/database"
	"pasarwarga-service-api/src/models"

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
	db := database.GetConnection()
	db.AutoMigrate(models.Article{})
	fmt.Println("Done")
}
