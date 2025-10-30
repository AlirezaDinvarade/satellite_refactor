package routes

import (
	"satellite/user/handlers"
	"satellite/user/models"

	"github.com/gofiber/fiber/v2"
)

func authRoutes(router fiber.Router) {
	redisAdaptor := &models.RedisAdaptor{Client: models.RedisDB}
	
	authHandler := handlers.NewAuthHandler(models.DB, redisAdaptor)
	auth := router.Group("/auth")

	auth.Post("/send-otp", authHandler.SendOTPHandler)
	auth.Post("/login-otp", authHandler.LoginOTPHandler)
	auth.Post("/set-password", authHandler.SetPasswordHandler)
}
