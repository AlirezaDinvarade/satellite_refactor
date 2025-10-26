package main

import (
	"log"
	"satellite/user/models"
	"satellite/user/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	var ()
	app := fiber.New()
	routes.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error in loading file .env")
	}
	models.CheckDatabaseConnection()
	models.CreateTables()
}
