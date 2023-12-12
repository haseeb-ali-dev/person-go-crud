package main

import (
	"go-crud/handlers"
	"log"
	"net/http"
)

// Handler Aliases
var (
	healthHandler = handlers.HealthHandler
	personHandler = handlers.PersonHandler
)

func main() {
	http.HandleFunc("/", healthHandler)
	http.HandleFunc("/person", personHandler)

	log.Println("Server is starting at :4444")

	err := http.ListenAndServe(":4444", nil)

	if err != nil {
		log.Fatal("Error starting server: ", err)
	}

	log.Println("Server is stopped")
}
