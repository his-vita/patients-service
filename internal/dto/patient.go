package dto

import (
	"cloud.google.com/go/civil"
	"github.com/google/uuid"
)

type Patient struct {
	ID         *uuid.UUID `json:"id"`
	FirstName  string     `json:"first_name"`
	LastName   string     `json:"last_name"`
	MiddleName *string    `json:"middle_name"`
	BirthDate  civil.Date `json:"birth_date"`
	Gender     *bool      `json:"gender"`
}

type PatientFull struct {
	Patient *Patient `json:"patient"`
	Contact *Contact `json:"contact"`
}

type PatientResponse struct {
	ID         *uuid.UUID
	FirstName  string
	LastName   string
	MiddleName *string
	BirthDate  civil.Date
	Gender     *bool
	Contact    *Contact
}
