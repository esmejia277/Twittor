package middleware

import (
	"net/http"

	"github.com/esmejia277/twittor/app/db"
)

func CheckDataBase(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if db.IsDBConnected() == 0 {
			http.Error(w, "Lost connection with db", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
