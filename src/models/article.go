package models

import "gorm.io/gorm"

type Article struct {
	ID         int      `json:"id"`
	Title      string   `json:"title"`
	CategoryID int      `json:"category_id"`
	Category   Category `json:"category"`
	Content    string   `json:"content"`
	gorm.Model
}
