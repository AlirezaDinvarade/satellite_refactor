package routes

import (
	"satellite/user/handlers"
	"satellite/user/models"

	"github.com/gofiber/fiber/v2"
)

func userRoutes(router fiber.Router) {
	userHander := handlers.NewUserHandler(models.DB)
	users := router.Group("/users")

	users.Post("/create", userHander.HandleCreateUser)
}
