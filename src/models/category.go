package models

import "gorm.io/gorm"

type Category struct {
	ID           int    `json:"id"`
	CategoryName string `json:"category_name"`
	CategorySlug string `json:"category_slug"`
	gorm.Model
	Articles []Article `gorm:"foreignKey:CategoryID" json:"articles"`
}
