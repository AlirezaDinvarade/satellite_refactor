package middlewares

import (
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "mysecrettoken" {
			http.Error(w, "bad", http.StatusInternalServerError)
		}
		next.ServeHTTP(w, r)
	})
}
