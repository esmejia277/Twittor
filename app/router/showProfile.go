package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/esmejia277/twittor/app/db"
)

func ShowProfile(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	fmt.Println(id)
	if len(id) < 1 {
		http.Error(w, "Send id parameter", http.StatusBadRequest)
		return
	}
	profile, err := db.SearchProfile(id)
	if err != nil {
		http.Error(w, "Something went wrong"+err.Error(), 400)
		return
	}
	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)

}
