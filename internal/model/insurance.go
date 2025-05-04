package model

import (
	"cloud.google.com/go/civil"
	"github.com/google/uuid"
)

type Insurance struct {
	ID                 *uuid.UUID  `json:"id,omitempty"`
	PatientID          *uuid.UUID  `json:"patient_id,omitempty"`
	Number             *string     `json:"number"`
	IssueDate          *civil.Date `json:"issue_date,omitempty"`
	ExpiryDate         *civil.Date `json:"expiry_date,omitempty"`
	Type               *int        `json:"type"`
	Main               *bool       `json:"main"`
	InsuranceCompanyID *int        `json:"insurance_company_id"`
}
