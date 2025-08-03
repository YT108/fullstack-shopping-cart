package controllers

import (
	"net/http"
	"shopping_cart/models"

	"github.com/gin-gonic/gin"
)

func AddToCart(c *gin.Context) {
	var req struct {
		UserID uint `json:"user_id"`
		ItemID uint `json:"item_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var cart models.Cart
	if err := DB.Where("user_id = ?", req.UserID).First(&cart).Error; err != nil {
		cart = models.Cart{UserID: req.UserID}
		DB.Create(&cart)
	}

	cartItem := models.CartItem{CartID: cart.ID, ItemID: req.ItemID}
	DB.Create(&cartItem)

	c.JSON(http.StatusOK, gin.H{"message": "Item added to cart"})
}

func ListCarts(c *gin.Context) {
	userId := c.Query("user_id")
	var cart models.Cart
	if err := DB.Preload("Items.Item").Where("user_id = ?", userId).First(&cart).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		return
	}
	c.JSON(http.StatusOK, cart)
}

func ClearCart(c *gin.Context) {
	userID := c.Param("user_id")

	// Find the cart belonging to this user
	var cart models.Cart
	if err := DB.Where("user_id = ?", userID).First(&cart).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		return
	}

	// Delete all items in the cart
	if err := DB.Where("cart_id = ?", cart.ID).Delete(&models.CartItem{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to clear cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cart cleared successfully"})
}
