package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"next_device/backend/db"
	"next_device/backend/di"
	"next_device/backend/routes"
)

func main() {
	app := fiber.New()

	dataBase, err := db.SetupDb()
	if err != nil {
		panic("Нет соединения с базой данных")
	}

	db.Migrate(dataBase)

	productController := di.InitDependencies(dataBase)
	routes.SetupRoutes(app, productController)

	server := app.Listen("127.0.0.1:5000")
	if server != nil {
		panic("Ошибка при запуске сервера")
	}

	fmt.Print("Сервер запущен")
}
