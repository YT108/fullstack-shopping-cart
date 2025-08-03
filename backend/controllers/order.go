package controllers

import (
	"net/http"
	"shopping_cart/models"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var req struct {
		UserID uint `json:"user_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var cart models.Cart
	if result := DB.Where("user_id = ?", req.UserID).First(&cart); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		return
	}

	order := models.Order{UserID: req.UserID, CartID: cart.ID}
	DB.Create(&order)

	c.JSON(http.StatusOK, gin.H{"message": "Order placed", "order_id": order.ID})
}

func ListOrders(c *gin.Context) {
	var orders []models.Order
	DB.Find(&orders)
	c.JSON(http.StatusOK, orders)
}
