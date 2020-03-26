package auth

import (
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Auth-Token")

		if isTokenValid(token) {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Access denied", http.StatusForbidden)
		}
	})
}

func isTokenValid(token string) bool {
	// @todo: implement token validation
	return true
}
