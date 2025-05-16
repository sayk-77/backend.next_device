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

	brandRep := repository.NewBrandRepository(db)
	brandService := service.NewBrandBrand(brandRep)
	brandController := controllers.NewBrandController(brandService)

	categoryRep := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRep)
	categoryController := controllers.NewCategoryController(categoryService)

	productRep := repository.NewProductRepository(db)
	productService := service.NewProductService(productRep, productImageService, brandRep, categoryRep)
	productController := controllers.NewProductController(productService)

	productDetailsRep := repository.NewProductDetailsRepository(db)
	productDetailsService := service.NewProductDetailsService(productDetailsRep)
	productDetailsController := controllers.NewProductDetailsController(productDetailsService)

	userRep := repository.NewUserRepository(db)
	userService := service.NewUserService(userRep)
	userController := controllers.NewUserController(userService)

	cartRep := repository.NewCartRepository(db)
	cartService := service.NewCartService(cartRep)
	cartController := controllers.NewCartController(cartService)

	orderRepository := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepository)
	orderController := controllers.NewOrderController(orderService)

	paymentRepo := repository.NewPaymentRepository(db)
	paymentService := service.NewPaymentService(paymentRepo)
	paymentController := controllers.NewPaymentController(paymentService, orderService)

	pushRepo := repository.NewPushRepository(db)
	pushService := service.NewNotificationService(pushRepo)
	pushController := controllers.NewNotificationController(pushService)

	reviewRepo := repository.NewReviewRepository(db)
	reviewService := service.NewReviewService(reviewRepo)
	reviewController := controllers.NewReviewController(reviewService, pushService, orderService)

	routes.SetupRoutes(app, productController, brandController, categoryController, productDetailsController, userController,
		cartController, paymentController, orderController, reviewController, pushController)
}
