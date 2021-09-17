package article

import (
	"fmt"
	"pasarwarga-service-api/database"
	"pasarwarga-service-api/src/models"

	"github.com/gosimple/slug"
)

// List all article with filtering support
func List(title string, category_name string) ([]models.Article, error) {
	// conn
	db := database.GetConnection()

	// prepare
	var articles []models.Article

	// query and parse
	if category_name != "" {
		db = db.Where("category_id IN ?", categorySearch(category_name))
	} else if title != "" {
		db = db.Where("title LIKE ?", fmt.Sprintf("%%%s%%", title))
	}
	result := db.Find(&articles)

	return articles, result.Error
}

func categorySearch(name_to_search string) []int {
	// db conn
	db := database.GetConnection()

	// prepare
	var categories []models.Category
	returnDump := []int{}

	// query and parse
	db.Where("category_name LIKE ?", fmt.Sprintf("%%%s%%", name_to_search)).Find(&categories)
	for _, category := range categories {
		returnDump = append(returnDump, category.ID)
	}
	return returnDump
}

// Get article by its' id
func Get(article_id int) (models.Article, error) {
	// conn
	db := database.GetConnection()

	var article models.Article
	db.First(&article, article_id)
	return article, db.Error
}

// Store new article
func Store(title string, category_id int, content string) (models.Article, error) {
	// conn
	db := database.GetConnection()

	// prepare
	article := models.Article{
		Title:      title,
		CategoryID: category_id,
		Content:    content,
	}

	result := db.Create(&article)
	return article, result.Error
}

// Update article entity
func Update(article_id int, title string, category_id int, content string) (models.Article, error) {
	// conn
	db := database.GetConnection()

	// update
	var article models.Article
	result := db.Model(&article).Where("id = ?", article_id).Updates(map[string]interface{}{
		"title":       title,
		"slug":        slug.Make(title),
		"category_id": category_id,
		"content":     content,
	})
	return article, result.Error
}

// Delete article
func Delete(article_id int) (models.Article, error) {
	db := database.GetConnection()

	var article models.Article
	result := db.Delete(&article, article_id)
	return article, result.Error
}
