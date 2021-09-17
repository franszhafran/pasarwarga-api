package models

import "gorm.io/gorm"

type Category struct {
	ID           int
	CategoryName string
	CategorySlug string
	gorm.Model
	Articles []Article `gorm:"foreignKey:CategoryID"`
}
