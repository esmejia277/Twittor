package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/esmejia277/twittor/app/middleware"
	routes "github.com/esmejia277/twittor/app/router"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middleware.CheckDataBase(routes.Register)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
