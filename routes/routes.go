package routes

import (
	"github.com/gofiber/fiber/v2"
	"next_device/backend/controllers"
)

func SetupRoutes(app *fiber.App, productController *controllers.ProductController) {
	api := app.Group("/api")

	api.Get("/products", productController.GetAllProducts)
	api.Get("/products/:id", productController.GetProductByID)
	api.Get("/catalog/category/:category", productController.GetProductsByCategory)
	api.Get("/catalog/discounts", productController.GetDiscountedProducts)
	api.Get("/catalog/new", productController.GetNewProducts)
	api.Post("/products", productController.CreateProduct)
	api.Put("/products/:id", productController.UpdateProduct)
	api.Delete("/products/:id", productController.DeleteProduct)
}
