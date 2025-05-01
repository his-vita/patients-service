package entity

import "github.com/google/uuid"

type Snils struct {
	PatientID *uuid.UUID `json:"patient_id"`
	Number    *string    `json:"number"`
}
