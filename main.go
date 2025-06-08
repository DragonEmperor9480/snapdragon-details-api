package main

import (
	"log"
	"net/http"
	"snapdragon-details-api/routes"
)

func main() {
	r := routes.RegisterRoutes()
	log.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
