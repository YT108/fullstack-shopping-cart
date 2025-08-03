package models

import "gorm.io/gorm"

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string
	Token    string `gorm:"-"`
}

type Item struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Price float64
}

type Cart struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint
	Items  []CartItem `gorm:"foreignKey:CartID"`
}

type CartItem struct {
	ID     uint `gorm:"primaryKey"`
	CartID uint
	ItemID uint
	Item   Item `gorm:"foreignKey:ItemID"` // ✅ Preload the full item
}

type Order struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint
	CartID uint
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Item{}, &Cart{}, &CartItem{}, &Order{})
}
