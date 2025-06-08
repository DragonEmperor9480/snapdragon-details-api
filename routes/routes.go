package routes

import (
	"snapdragon-details-api/controllers"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/processors", controllers.GetAllProcessors).Methods("GET")
	return r
}
