package routes

import (
	"shopping_cart/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	controllers.InitModels(db)

	r.POST("/users", controllers.CreateUser)
	r.GET("/users", controllers.ListUsers)
	r.POST("/users/login", controllers.LoginUser)

	r.POST("/items", controllers.CreateItem)
	r.GET("/items", controllers.ListItems)

	r.POST("/carts", controllers.AddToCart)
	r.GET("/carts", controllers.ListCarts)

	r.DELETE("/carts/user/:user_id", controllers.ClearCart)

	r.POST("/orders", controllers.CreateOrder)
	r.GET("/orders", controllers.ListOrders)
}
