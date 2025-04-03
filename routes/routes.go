package routes

import (
	"github.com/gorilla/mux"
	"enormous-air-portal/controllers"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/bookings", controllers.CreateBooking).Methods("POST")
	return router
}
