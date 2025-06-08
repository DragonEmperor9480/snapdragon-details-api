package controllers

import (
	"encoding/json"
	"net/http"
	"snapdragon-details-api/models"
	"snapdragon-details-api/utils"

	"golang.org/x/crypto/bcrypt"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	utils.DB.Create(&user)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User Created Sucessfully"})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var inputUser models.User
	var dbUser models.User

	json.NewDecoder(r.Body).Decode(&inputUser)

	utils.DB.Where("email=?", inputUser.Email).First(&dbUser)

	if dbUser.ID == 0 {
		http.Error(w, "User not Found", http.StatusUnauthorized)
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(inputUser.Password))
	if err != nil {
		http.Error(w, "Invalid Password", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Login Sucessful"})

}
