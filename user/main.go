package main

import (
	"fmt"
	"log"
	"net/http"
	"satellite/user/middlewares"
	"satellite/user/routes"
	"satellite/user/stores"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	router := mux.NewRouter()
	router.Use(middlewares.LoggingMiddleware)
	router.Use(middlewares.AuthMiddleware)

	routes.SetupRoutes(router)
	fmt.Println("Server is running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error in loading file .env")
	}
	stores.CheckDatabaseConnection()
	stores.CreateTables()
}
