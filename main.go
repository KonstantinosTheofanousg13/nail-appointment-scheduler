package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/KonstantinosTheofanousg13/nail-appointment-scheduler/database"
	"github.com/KonstantinosTheofanousg13/nail-appointment-scheduler/handlers"
)

func main() {
	database.Connect()

	http.HandleFunc("POST /appointments", handlers.CreateAppointment)
	http.HandleFunc("GET /appointments/", handlers.GetAppointments)
	http.HandleFunc("DELETE /appointments/{id}", handlers.DeleteAppointment)
	http.HandleFunc("PUT /appointments/{id}", handlers.UpdateAppointment)

	fmt.Println("Server is running on port 8080...")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
