package router

import (
	"encoding/json"
	"net/http"

	"github.com/esmejia277/twittor/app/db"
	"github.com/esmejia277/twittor/app/models"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error decoding json payload"+err.Error(), 500)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email is required", 400)
	}
	if len(t.Password) < 6 {
		http.Error(w, "Password has to be al least 7 characters length", 400)
	}
	_, userFound, _ := db.CheckIfUserExists(t.Email)
	if userFound {
		http.Error(w, "User already exists", 400)
		return
	}

	_, status, err := db.InsertIntoDatabase(t)
	if err != nil {
		http.Error(w, "Error trying to insert into database"+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "Error trying to insert into database"+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
