package db

import (
	"fmt"
	"gorm.io/gorm"
	"next_device/backend/models"
)

func Migrate(database *gorm.DB) {
	err := database.AutoMigrate(models.User{}, models.Cart{}, models.ProductVariant{},
		models.Products{}, models.Address{}, models.Brand{}, models.Category{}, models.Order{},
		models.OrderItem{}, models.Review{}, models.Payment{}, models.ProductImage{}, models.BrandBanners{},
		models.ProductDetails{})
	if err != nil {
		fmt.Println(err)
		return
	}
}
