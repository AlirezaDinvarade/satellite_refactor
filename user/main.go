package main

import (
	"log"
	"satellite/user/models"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	var ()
	app := fiber.New()
	app.Get("/")
	app.Listen(":3000")
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error in loading file .env")
	}
	models.CheckDatabaseConnection()
	models.CreateTables()
}
