package entity

import (
	"github.com/google/uuid"
)

type Contact struct {
	PatientID       *uuid.UUID `json:"patient_id"`
	PhoneNumber     *string    `json:"phone_number"`
	WorkPhoneNumber *string    `json:"work_phone_number"`
	Email           *string    `json:"email"`
}
