package main

import (
	"fmt"
	"log"
	"net/http"

	"enormous-air-portal/config"
	"enormous-air-portal/routes"
)

func main() {
	db := config.InitDB()
	defer db.Close()

	router := routes.InitRoutes()

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
