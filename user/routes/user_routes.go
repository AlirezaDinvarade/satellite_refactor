package routes

import (
	"github.com/gorilla/mux"
)

func userRoutes(router *mux.Router) {
	users := router.PathPrefix("/users").Subrouter()
	users.HandleFunc("/create", userHander.HandleCreateUser).Methods("POST")
}
