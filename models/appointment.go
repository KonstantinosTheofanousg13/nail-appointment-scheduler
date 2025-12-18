package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Appointment struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Customer    string             `bson:"customer_name" json:"customer_name"`
	Service     string             `bson:"service_type" json:"service_type"`
	Date        time.Time          `bson:"appointment_date" json:"appointment_date"`
	PhoneNumber string             `bson:"phone" json:"phone"`
	Status      string             `bson:"status" json:"status"`
}
