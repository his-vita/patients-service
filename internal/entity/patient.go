package entity

import (
	"cloud.google.com/go/civil"
	"github.com/google/uuid"
)

type Patient struct {
	ID         uuid.UUID       `json:"id"`
	FirstName  string          `json:"first_name"`
	LastName   string          `json:"last_name"`
	MiddleName *string         `json:"middle_name"`
	BirthDate  civil.Date      `json:"birth_date"`
	Gender     *bool           `json:"gender"`
	CreatedTS  *civil.DateTime `json:"created_ts"`
	CreatedBy  *string         `json:"created_by"`
	UpdatedTS  *civil.DateTime `json:"updated_ts"`
	UpdatedBy  *string         `json:"updated_by"`
	DeletedTS  *civil.DateTime `json:"deleted_ts"`
	DeletedBy  *string         `json:"deleted_by"`
	Version    int             `json:"version"`
	Contact    Contact         `json:"contact"`
}
