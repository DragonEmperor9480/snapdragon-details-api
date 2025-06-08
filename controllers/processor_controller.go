package controllers

import (
	"encoding/json"
	"net/http"
	"snapdragon-details-api/models"
)

var processors = []models.Processor{
	{ID: 1, Name: "Snapdragon 888", Cores: 8, Frequency: "2.84 GHz"},
	{ID: 2, Name: "Snapdragon 8 Gen 1", Cores: 8, Frequency: "3.0 GHz"},
}

func GetAllProcessors(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(processors)
}