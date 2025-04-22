package dto

import (
	"cloud.google.com/go/civil"
	"github.com/google/uuid"
)

type PatientDTO struct {
	ID         uuid.UUID
	FirstName  string
	LastName   string
	MiddleName *string
	BirthDate  civil.Date
	Gender     *bool
	Contact    *ContactDTO
}
