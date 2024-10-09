package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"next_device/backend/db"
)

func main() {
	app := fiber.New()

	dataBase, err := db.SetupDb()
	if err != nil {
		panic("Нет соединения с базой данных")
	}

	db.Migrate(dataBase)

	server := app.Listen("127.0.0.1:5000")
	if server != nil {
		panic("Ошибка при запуске сервера")
	}

	fmt.Print("Сервер запущен")
}
