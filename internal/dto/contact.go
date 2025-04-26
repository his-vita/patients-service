package dto

import "github.com/google/uuid"

type Contact struct {
	ID              *uuid.UUID `json:"id"`
	PatientId       *uuid.UUID `json:"patient_id,omitempty"`
	WorkPhoneNumber *string    `json:"work_phone_number"`
	PhoneNumber     *string    `json:"phone_number"`
	Email           *string    `json:"email"`
}
