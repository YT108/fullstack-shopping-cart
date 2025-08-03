package controllers

import (
	"net/http"
	"shopping_cart/models"

	"github.com/gin-gonic/gin"
)

func CreateItem(c *gin.Context) {
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	DB.Create(&item)
	c.JSON(http.StatusOK, item)
}

func ListItems(c *gin.Context) {
	var items []models.Item
	DB.Find(&items)
	c.JSON(http.StatusOK, items)
}
