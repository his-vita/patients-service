package models

import "github.com/google/uuid"

type Patient struct {
	Id          uuid.UUID `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	MiddleName  string    `json:"middle_name"`
	BirthDate   string    `json:"birth_date"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
}
