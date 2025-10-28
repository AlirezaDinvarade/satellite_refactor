package main

import (
	"log"
	"satellite/user/handlers"
	"satellite/user/models"
	"satellite/user/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var config = fiber.Config{
	ErrorHandler: handlers.ErrorHandler,
}

func main() {
	var ()
	app := fiber.New(config)
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
	models.ConnectRedis()
}
