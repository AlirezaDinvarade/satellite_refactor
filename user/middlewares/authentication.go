package middlewares

import (
	"net/http"
	"satellite/user/handlers"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "mysecrettoken" {
			handlers.WriteJson(w, http.StatusForbidden, map[string]string{"detail": "You do not have proper authentication header"})
			return
		}
		next.ServeHTTP(w, r)
	})
}
