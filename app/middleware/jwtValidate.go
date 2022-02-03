package middleware

import (
	"net/http"

	"github.com/esmejia277/twittor/app/router"
)

func JWTValidate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := router.ProcessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Token error!"+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
