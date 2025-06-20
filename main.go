package main

import (
	"log"
	"net/http"
	"snapdragon-details-api/controllers"
	"snapdragon-details-api/models"
	"snapdragon-details-api/utils"

	"github.com/gorilla/mux"
)

func main() {
	utils.ConnectDB()

	utils.DB.AutoMigrate(&models.Processor{})

	r := mux.NewRouter()

	//APIs
	r.HandleFunc("/api/processors", controllers.CreateProcessor).Methods("POST")
	r.HandleFunc("/api/processors", controllers.GetAllProcessors).Methods("GET")
	r.HandleFunc("/api/processors/{id}", controllers.UpdateProcessor).Methods("PUT")
	r.HandleFunc("/api/processors/{id}", controllers.DeleteProcessor).Methods("DELETE")

	//login and Signup API
	r.HandleFunc("/api/signup", controllers.Signup).Methods("POST")
	r.HandleFunc("/api/login", controllers.Login).Methods("POST")

	log.Println("Server running on port http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
