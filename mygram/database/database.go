package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"final-project/models"
)

func SetupDB() *gorm.DB {

	DB_HOST := "127.0.0.1"
	DB_USER := "postgres"
	DB_PASSWORD := "postgres"
	DB_NAME := "mygram_db"
	DB_PORT := "5432"

	URL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_NAME, DB_PASSWORD)
	db, err := gorm.Open(postgres.Open(URL), &gorm.Config{})

	if err != nil {
		panic("Failed connect to database")
	}

	err = db.AutoMigrate(&models.User{}, &models.Photo{}, &models.Comment{}, &models.SocialMedia{})

	if err != nil {
		errorMessage := fmt.Sprintf("Failed to migrate models: %s", err.Error())
		panic(errorMessage)
	}

	fmt.Println("Database Connected")

	return db
}
