package routes

import (
	"satellite/user/handlers"
	"satellite/user/stores"

	"github.com/gorilla/mux"
)

var (
	userStore    = stores.NewPostgresUserStore(stores.DB)
	redisAdaptor = stores.NewRedisAdaptor()
	store = &stores.Store{
		User: userStore,
	}
	
	userHander  = handlers.NewUserHandler(store)
	authHandler = handlers.NewAuthHandler(stores.DB, redisAdaptor)
)

func SetupRoutes(router *mux.Router) {
	api := router.PathPrefix("/api").Subrouter()
	userRoutes(api)
	authRoutes(api)
}
