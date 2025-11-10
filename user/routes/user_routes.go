package routes

import (
	"github.com/gofiber/fiber/v2"
)

func userRoutes(router fiber.Router) {
	users := router.Group("/users")
	users.Post("/create", userHander.HandleCreateUser)
}
