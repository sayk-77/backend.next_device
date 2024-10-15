package di

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"next_device/backend/controllers"
	"next_device/backend/repository"
	"next_device/backend/routes"
	"next_device/backend/service"
)

func InitDependencies(app *fiber.App, db *gorm.DB) {
	productImageRep := repository.NewProductImageRepository(db)
	productImageService := service.NewProductImageService(productImageRep)

	productRep := repository.NewProductRepository(db)
	productService := service.NewProductService(productRep, productImageService)
	productController := controllers.NewProductController(productService)

	brandRep := repository.NewBrandRepository(db)
	brandService := service.NewBrandBrand(brandRep)
	brandController := controllers.NewBrandController(brandService)

	categoryRep := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRep)
	categoryController := controllers.NewCategoryController(categoryService)

	routes.SetupRoutes(app, productController, brandController, categoryController)
}
