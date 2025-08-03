// main.go
package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"shopping_cart/controllers"
	"shopping_cart/routes"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:@tcp(127.0.0.1:3306)/shopping_cart?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB: ", err)
	}
	DB = db
	controllers.InitModels(DB)
}

func main() {
	InitDB()

	r := gin.Default()

	// 🔐 Enable CORS to allow frontend (localhost:5173) to call backend (localhost:8080)
	r.Use(cors.Default())

	routes.RegisterRoutes(r, DB)

	// Start backend server
	r.Run(":8080")
}
