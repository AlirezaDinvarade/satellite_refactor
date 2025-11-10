package routes

import (
	"satellite/user/handlers"
	"satellite/user/stores"

	"github.com/gofiber/fiber/v2"
)

var (
	userStore = stores.NewPostgresUserStore(stores.DB)
	store     = &stores.Store{
		User: userStore,
	}
	userHander = handlers.NewUserHandler(store)
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	userRoutes(api)
	authRoutes(api)

}
