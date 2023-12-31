package main

import (
	"log"
	"net/http"
)

type application struct {}

func main() {
	// app config
	app := application{}

	// get application routes
	mux := app.routes()

	// print start message
	log.Println("Starting server on port 8080...")

	// start server
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}