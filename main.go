package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"next_device/backend/db"
	"next_device/backend/di"
	"next_device/backend/tools"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	dataBase, err := db.SetupDb()
	if err != nil {
		panic("Нет соединения с базой данных")
	}

	db.Migrate(dataBase)

	di.InitDependencies(app, dataBase)
	tools.GetImageProduct(app)
	tools.GetImageBrand(app)
	tools.GetImageBanner(app)
	tools.GetImageCategory(app)

	server := app.Listen("localhost:5000")
	if server != nil {
		panic("Ошибка при запуске сервера")
	}

	fmt.Print("Сервер запущен")
}
