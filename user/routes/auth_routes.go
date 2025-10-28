package routes

import (
	"satellite/user/handlers"
	"satellite/user/models"

	"github.com/gofiber/fiber/v2"
)

func authRoutes(router fiber.Router) {
	authHandler := handlers.NewAuthHandler(models.DB, models.RedisDB)
	auth := router.Group("/auth")

	auth.Post("/send-otp", authHandler.SendOTPHandler)
	auth.Post("/login-otp", authHandler.LoginOTP)
}
