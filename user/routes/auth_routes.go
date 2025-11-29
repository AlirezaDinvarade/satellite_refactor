package routes

import (
	"github.com/gorilla/mux"
)

func authRoutes(router *mux.Router) {

	auth := router.PathPrefix("/auth").Subrouter()

	auth.HandleFunc("/send-otp", authHandler.SendOTPHandler).Methods("POST")
	auth.HandleFunc("/login-otp", authHandler.LoginOTPHandler).Methods("POST")
}
