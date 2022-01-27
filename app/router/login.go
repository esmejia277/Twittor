package router

import (
	"encoding/json"
	"net/http"

	"github.com/esmejia277/twittor/app/db"
	"github.com/esmejia277/twittor/app/jwt"
	"github.com/esmejia277/twittor/app/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	var model models.User
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		http.Error(w, "User or password invalid: "+err.Error(), 400)
	}
	if len(model.Email) == 0 {
		http.Error(w, "Email required", 400)
		return
	}
	user, exists := db.TryLogin(model.Email, model.Password)

	if !exists {
		http.Error(w, "User doest not exists "+err.Error(), 400)
	}

	jwtToken, err := jwt.GenerateJWT(user)
	if err != nil {
		http.Error(w, "Authentication error "+err.Error(), 400)
		return
	}

	response := models.LoginResponse{
		Token: jwtToken,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

}
