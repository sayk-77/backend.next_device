package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDb() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("ошибка загрузки переменных окружения: %v", err)
	}

	dbUser := os.Getenv("USERNAME_DB")
	dbPassword := os.Getenv("PASSWORD_DB")
	dbPort := os.Getenv("PORT")
	dbHost := os.Getenv("HOST")
	dbName := os.Getenv("DB_NAME")
	dbIsSslMode := os.Getenv("IS_SSL_MODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", dbHost, dbUser, dbPassword, dbName, dbPort, dbIsSslMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
		return nil, err
	}

	return db, nil
}
