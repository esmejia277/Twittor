package main

import (
	"log"

	"github.com/esmejia277/twittor/app/db"
	"github.com/esmejia277/twittor/app/handlers"
)

func main() {

	if db.IsDBConnected() == 1 {
		log.Print("Connected to DB succesfully")
	} else {
		log.Fatal("No conneted to database")
	}

	handlers.Handlers()
}
