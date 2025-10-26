package routes

import (
	"satellite/user/handlers"
	"satellite/user/models"

	"github.com/gofiber/fiber/v2"
)

func authRoutes(router fiber.Router) {
	userHander := handlers.NewUserHandler(models.DB)
	users := router.Group("/auth")

	users.Post("/login", userHander.HandleCreateUser)
}
