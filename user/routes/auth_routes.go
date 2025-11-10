package routes

import (
	"satellite/user/handlers"
	"satellite/user/stores"

	"github.com/gofiber/fiber/v2"
)

func authRoutes(router fiber.Router) {
	redisAdaptor := &stores.RedisAdaptor{Client: stores.RedisDB}

	authHandler := handlers.NewAuthHandler(stores.DB, redisAdaptor)
	auth := router.Group("/auth")

	auth.Post("/send-otp", authHandler.SendOTPHandler)
	auth.Post("/login-otp", authHandler.LoginOTPHandler)
	// auth.Post("/set-password", authHandler.SetPasswordHandler)
}
