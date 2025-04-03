package controllers

import (
	"encoding/json"
	"net/http"
	"enormous-air-portal/models"
)

func CreateBooking(w http.ResponseWriter, r *http.Request) {
	var booking models.Booking
	json.NewDecoder(r.Body).Decode(&booking)

	err := models.AddBooking(booking)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Booking created successfully")
}
