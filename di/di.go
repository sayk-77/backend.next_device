package di

import (
	"gorm.io/gorm"
	"next_device/backend/controllers"
	"next_device/backend/repository"
	"next_device/backend/service"
)

func InitDependencies(db *gorm.DB) *controllers.ProductController {
	productImageRep := repository.NewProductImageRepository(db)
	productImageService := service.NewProductImageService(productImageRep)
	productRep := repository.NewProductRepository(db)
	productService := service.NewProductService(productRep, productImageService)
	return controllers.NewProductController(productService)
}
