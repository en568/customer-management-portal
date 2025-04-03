package controllers

import (
	"encoding/json"
	"net/http"
	"enormous-air-portal/models"
	"enormous-air-portal/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	dbUser, err := models.FindUserByEmail(user.Email)
	if err != nil || !utils.ComparePassword(dbUser.Password, user.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, _ := utils.GenerateJWT(dbUser.Email, dbUser.Role)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
