package routes

import (
	"github.com/gofiber/fiber/v2"
	"next_device/backend/controllers"
	"next_device/backend/tools"
)

func SetupRoutes(app *fiber.App, productController *controllers.ProductController,
	brandController *controllers.BrandController, categoryController *controllers.CategoryController,
	productDetailsController *controllers.ProductDetailsController, userController *controllers.UserController,
	cartController *controllers.CartController, paymentController *controllers.PaymentController,
	orderController *controllers.OrderController, reviewController *controllers.ReviewController,
	pushController *controllers.NotificationController) {
	api := app.Group("/api")

	api.Post("/subscribe", pushController.Subscribe)
	api.Post("/push-not", pushController.SendNotification)
	api.Get("/products", productController.GetAllProducts)
	api.Put("/products", productController.UpdateProduct)
	api.Get("/products/category", productController.GetProductsByBrandAndCategory)
	api.Get("/search", productController.SearchProduct)
	api.Post("/product/laptop/query", productController.GetFilteredLaptops)
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
	api.Post("/product/images", productController.CreateImageProduct)
	api.Get("/product/details/:id", productDetailsController.GetProductDetails)
	api.Post("/product/details", productDetailsController.CreateProductDetails)
	api.Put("/product/details/:id", productDetailsController.UpdateProductDetails)
	api.Delete("/product/details/:id", productDetailsController.DeleteProductDetails)
	api.Post("/register", userController.Register)
	api.Post("/login", userController.Login)
	api.Put("/user/update", tools.JWTMiddleware, userController.SaveUserInfo)
	api.Put("/user/password", tools.JWTMiddleware, userController.ChangePassword)
	api.Post("/user/address", tools.JWTMiddleware, userController.AddNewAddress)
	api.Delete("/user/address/:id", tools.JWTMiddleware, userController.DeleteAddress)
	api.Post("/cart/add", tools.JWTMiddleware, cartController.AddItem)
	api.Delete("/cart/remove/:productId", tools.JWTMiddleware, cartController.RemoveItem)
	api.Get("/cart", tools.JWTMiddleware, cartController.GetCartItems)
	api.Get("/user", tools.JWTMiddleware, userController.GetUserById)
	api.Post("/payment", tools.JWTMiddleware, paymentController.HandlePaymentIntent)
	api.Post("/payment/status", tools.JWTMiddleware, paymentController.HandlePaymentStatus)
	api.Get("/order/all", tools.JWTMiddleware, orderController.GetAllOrders)
	api.Get("/order/:id", tools.JWTMiddleware, orderController.GetOrderById)
	api.Get("/review/all", tools.JWTMiddleware, reviewController.GetAllReviews)
	api.Post("/order/status", reviewController.ChangeStatus)
	api.Get("/review/product/:id", reviewController.GetReviewForProduct)
	api.Get("/review/:id", tools.JWTMiddleware, reviewController.GetReviewById)
	api.Put("/review/:id", tools.JWTMiddleware, reviewController.PublishReview)
	api.Post("/review", tools.JWTMiddleware, reviewController.CreateReview)
	api.Delete("/review/:id", tools.JWTMiddleware, reviewController.DeleteReview)
}
