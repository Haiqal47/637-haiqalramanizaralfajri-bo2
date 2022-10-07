package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"ass-02/models"
)

func SetupDB() *gorm.DB {
	err := godotenv.Load(".env")

	if err != nil {
		panic("Failed to load env file")
	}

	DB_HOST := os.Getenv("DB_HOST")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")

	URL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_NAME, DB_PASSWORD)
	db, err := gorm.Open(postgres.Open(URL), &gorm.Config{})

	if err != nil {
		panic("Failed connect to database")
	}

	err = db.AutoMigrate(&models.Order{}, &models.Item{})

	if err != nil {
		errorMessage := fmt.Sprintf("Failed to migrate models: %s", err.Error())
		panic(errorMessage)
	}

	fmt.Println("Database Connected")

	return db
}
