package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/KonstantinosTheofanousg13/nail-appointment-scheduler/database"
	"github.com/KonstantinosTheofanousg13/nail-appointment-scheduler/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateAppointment(responseWriter http.ResponseWriter, incomingRequest *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")

	var appointment models.Appointment
	if err := json.NewDecoder(incomingRequest.Body).Decode(&appointment); err != nil {
		http.Error(responseWriter, err.Error(), http.StatusBadRequest)
		return
	}

	if appointment.Date.IsZero() {
		appointment.Date = time.Now()
	}
	appointment.Status = "Pending"

	fiveSecondTimer, stopTimer := database.GetContext()
	defer stopTimer()

	collection := database.Client.Database("nail_scheduler").Collection("appointments")
	result, err := collection.InsertOne(fiveSecondTimer, appointment)
	if err != nil {
		http.Error(responseWriter, "Failed to save appointment", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(responseWriter).Encode(result); err != nil {
		log.Println("Error sending response:", err)
	}
}

func GetAppointments(responseWriter http.ResponseWriter, incomingRequest *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")

	collection := database.Client.Database("nail_scheduler").Collection("appointments")

	fiveSecondTimer, stopTimer := database.GetContext()
	defer stopTimer()

	cursor, err := collection.Find(fiveSecondTimer, bson.M{})
	if err != nil {
		http.Error(responseWriter, "Failed to fetch appointments", http.StatusInternalServerError)
		return
	}

	defer func() {
		if err := cursor.Close(fiveSecondTimer); err != nil {
			log.Println("Error closing cursor:", err)
		}
	}()

	var appointments []models.Appointment
	if err = cursor.All(fiveSecondTimer, &appointments); err != nil {
		http.Error(responseWriter, "Failed to parse appointments", http.StatusInternalServerError)
		return
	}

	if appointments == nil {
		appointments = []models.Appointment{}
	}

	if err = json.NewEncoder(responseWriter).Encode(appointments); err != nil {
		log.Println("Error sending response:", err)
	}
}

func DeleteAppointment(responseWriter http.ResponseWriter, incomingRequest *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")

	idParam := incomingRequest.PathValue("id")
	objectId, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		http.Error(responseWriter, "Invalid ID format", http.StatusBadRequest)
		return
	}

	collection := database.Client.Database("nail_scheduler").Collection("appointments")

	fiveSecondTimer, stopTimer := database.GetContext()
	defer stopTimer()

	result, err := collection.DeleteOne(fiveSecondTimer, bson.M{"_id": objectId})
	if err != nil {
		http.Error(responseWriter, "Failed to delete appointment", http.StatusInternalServerError)
		return
	}

	if result.DeletedCount == 0 {
		http.Error(responseWriter, "Appointment not found", http.StatusNotFound)
		return
	}

	err = json.NewEncoder(responseWriter).Encode(map[string]string{"message": "Appointment deleted successfully"})
	if err != nil {
		log.Println("Error sending response:", err)
	}
}

func UpdateAppointment(responseWriter http.ResponseWriter, incomingRequest *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")

	idParam := incomingRequest.PathValue("id")
	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		http.Error(responseWriter, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var updatedData models.Appointment
	if err := json.NewDecoder(incomingRequest.Body).Decode(&updatedData); err != nil {
		http.Error(responseWriter, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	collection := database.Client.Database("nail_scheduler").Collection("appointments")
	fiveSecondTimer, stopTimer := database.GetContext()
	defer stopTimer()

	update := bson.M{
		"$set": bson.M{
			"appointment_date": updatedData.Date,
			"service_type":     updatedData.Service,
			"status":           updatedData.Status,
		},
	}

	result, err := collection.UpdateOne(fiveSecondTimer, bson.M{"_id": objectID}, update)
	if err != nil {
		http.Error(responseWriter, "Failed to update appointment", http.StatusInternalServerError)
		return
	}

	if result.MatchedCount == 0 {
		http.Error(responseWriter, "Appointment not found", http.StatusNotFound)
		return
	}

	err = json.NewEncoder(responseWriter).Encode(map[string]string{"message": "Appointment updated successfully"})
	if err != nil {
		log.Println("Error sending response:", err)
	}

}
