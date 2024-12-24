package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var (
	db *gorm.DB
)

func Connect() {
	// Завантаження .env файлу
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// Формування DSN з використанням змінних середовища
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
		os.Getenv("DB_TIMEZONE"))

	// Встановлення підключення до бази даних
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db = d
	fmt.Printf("Database connected!!!")
}

func GetDB() *gorm.DB {
	return db
}
