package main

import (
	"hello-jwt/database"
	"hello-jwt/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
