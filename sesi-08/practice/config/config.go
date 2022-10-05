package config

import (
	"practice/structs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBInit() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=gorm port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(structs.Person{})
	return db
}
