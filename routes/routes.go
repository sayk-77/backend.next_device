package routes

import (
	"github.com/gofiber/fiber/v2"
	"next_device/backend/controllers"
)

func SetupRoutes(app *fiber.App, productController *controllers.ProductController,
	brandController *controllers.BrandController, categoryController *controllers.CategoryController,
	productDetailsController *controllers.ProductDetailsController) {
	api := app.Group("/api")

	api.Get("/products", productController.GetAllProducts)
	api.Get("/products/category", productController.GetProductsByBrandAndCategory)
	api.Get("/search", productController.SearchProduct)
	api.Post("/product/:category/query", productController.GetFilteredProducts)
	api.Get("/products/:param", productController.GetProductByIdOrName)
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
	api.Get("/brands/:id/category", brandController.GetCategoriesByBrand)
	api.Get("/categories", categoryController.GetAllCategories)
	api.Get("/categories/count", categoryController.GetAllCategoryAndCount)
	api.Get("/categories/:param", categoryController.GetCategoryByIdOrName)
	api.Post("/categories", categoryController.CreateCategory)
	api.Put("/categories/:id", categoryController.UpdateCategory)
	api.Delete("/categories/:id", categoryController.DeleteCategory)
	api.Get("/product/details/:id", productDetailsController.GetProductDetails)
	api.Post("/product/details", productDetailsController.CreateProductDetails)
	api.Put("/product/details/:id", productDetailsController.UpdateProductDetails)
	api.Delete("/product/details/:id", productDetailsController.DeleteProductDetails)
}
