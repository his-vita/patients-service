package entity

import (
	"github.com/google/uuid"
)

type Contact struct {
	ID          uuid.UUID `json:"id"`
	PatientID   uuid.UUID `json:"patient_id"`
	PhoneNumber *string   `json:"phone_number"`
	Email       *string   `json:"email"`
	Main        bool      `json:"main"`
}
