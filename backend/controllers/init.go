package controllers

import (
	"shopping_cart/models"

	"gorm.io/gorm"
)

var DB *gorm.DB

func InitModels(db *gorm.DB) {
	DB = db
	models.AutoMigrate(DB)
}
