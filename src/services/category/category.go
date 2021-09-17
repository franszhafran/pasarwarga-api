package category

import (
	"pasarwarga-service-api/database"
	"pasarwarga-service-api/src/models"

	"github.com/gosimple/slug"
)

func List() ([]models.Category, error) {
	db := database.GetConnection()

	var categories []models.Category
	db.Find(&categories)
	return categories, db.Error
}

func Get(category_id int) (models.Category, error) {
	db := database.GetConnection()

	var category models.Category
	db.First(&category, category_id)
	return category, db.Error
}

func Store(name string) (models.Category, error) {
	db := database.GetConnection()

	category := models.Category{
		CategoryName: name,
		CategorySlug: slug.Make(name),
	}
	result := db.Create(&category)
	return category, result.Error
}

func Update(category_id int, name string) (models.Category, error) {
	db := database.GetConnection()

	var category models.Category
	result := db.Model(&category).Where("id = ?", category_id).Update("category_name", name)
	return category, result.Error
}

func Delete(category_id int) (models.Category, error) {
	db := database.GetConnection()

	var category models.Category
	result := db.Delete(&category, category_id)
	return category, result.Error
}
