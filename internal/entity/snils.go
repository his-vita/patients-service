package entity

import "github.com/google/uuid"

type Snils struct {
	ID        *uuid.UUID `json:"id"`
	PatientId *uuid.UUID `json:"patient_id"`
	Number    *string    `json:"number"`
}
