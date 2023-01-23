package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jinheehanaaa/go-blog-backend/database"
	"github.com/joho/godotenv"
)

func main() {
	database.Connect()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env files")
	}
	port := os.Getenv("PORT")
	app := fiber.New()
	app.Listen(":" + port)
}
