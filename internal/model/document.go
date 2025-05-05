package model

import (
	"cloud.google.com/go/civil"
	"github.com/google/uuid"
)

type Document struct {
	ID                *uuid.UUID  `json:"id"`
	Series            *string     `json:"series"`
	Number            *string     `json:"number"`
	DepartmentCode    *string     `json:"department_code"`
	IssueDate         *civil.Date `json:"issue_date"`
	ExpiryDate        *civil.Date `json:"expiry_date"`
	Main              *bool       `json:"main"`
	PatientID         *uuid.UUID  `json:"patient_id,omitempty"`
	DocumentTypeID    *int        `json:"document_type_id"`
	DocumentCompanyID *int        `json:"document_company_id"`
}
