package controllers

import (
	"encoding/json"
	"net/http"
	"snapdragon-details-api/models"
	"snapdragon-details-api/utils"

	"github.com/gorilla/mux"
)

func CreateProcessor(w http.ResponseWriter, r *http.Request) {
	var processor models.Processor

	if err := json.NewDecoder(r.Body).Decode(&processor); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	result := utils.DB.Create(&processor)
	if result.Error != nil {
		http.Error(w, "Failed to save processor", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(processor)
}

func GetAllProcessors(w http.ResponseWriter, r *http.Request) {
	var processors []models.Processor

	result := utils.DB.Find(&processors)
	if result.Error != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(processors)

}

func UpdateProcessor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]

	var processor models.Processor

	if err := utils.DB.First(&processor, id).Error; err != nil {
		http.Error(w, "Processor not found", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&processor); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := utils.DB.Save(&processor).Error; err != nil {
		http.Error(w, "Failed to update processor", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(processor)
}
