package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting server in localhost and port 8000")
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8000", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello from go package")
}
