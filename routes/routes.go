package routes

import (
	"github.com/gofiber/fiber/v2"
	"next_device/backend/controllers"
)

func SetupRoutes(app *fiber.App, productController *controllers.ProductController, brandController *controllers.BrandController) {
	api := app.Group("/api")

	api.Get("/products", productController.GetAllProducts)
	api.Get("/products/:id", productController.GetProductByID)
	api.Get("/catalog/discounts", productController.GetDiscountedProducts)
	api.Get("/catalog/new", productController.GetNewProducts)
	api.Get("/catalog/:category", productController.GetProductsByCategory)
	api.Post("/products", productController.CreateProduct)
	api.Put("/products/:id", productController.UpdateProduct)
	api.Delete("/products/:id", productController.DeleteProduct)
	api.Get("/brands", brandController.GetAllBrand)
	api.Get("/brands/:param", brandController.GetBrandByNameOrId)
	api.Post("/brands", brandController.CreateBrand)
	api.Put("/brands/:id", brandController.UpdateBrand)
	api.Delete("/brands/:id", brandController.DeleteBrand)
}
