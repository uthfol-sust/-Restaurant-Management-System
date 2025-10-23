package middleware

import (
	"net/http"
	"restaurant-system/pkg/utils"
)

func AuthMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie , err := r.Cookie("jwt")

		if err != nil {
			http.Error(w, "Unauthorized: missing cookie", http.StatusUnauthorized)
			return
		}

		_ , err = utils.VerifyByJWT(cookie.Value)
		if err != nil {
			http.Error(w, "Unauthorized: invalid or expired token", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w,r)
	})
}
