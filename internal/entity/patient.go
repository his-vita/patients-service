package entity

import (
	"database/sql"

	"cloud.google.com/go/civil"
	"github.com/google/uuid"
)

type Patient struct {
	Id          uuid.UUID      `json:"id"`
	FirstName   string         `json:"first_name"`
	LastName    string         `json:"last_name"`
	MiddleName  sql.NullString `json:"middle_name"`
	BirthDate   civil.Date     `json:"birth_date"`
	Gender      sql.NullBool   `json:"gender"`
	PhoneNumber sql.NullString `json:"phone_number"`
	Email       sql.NullString `json:"email"`
	CreatedTS   sql.NullTime   `json:"created_ts"`
	CreatedBy   sql.NullString `json:"created_by"`
	UpdatedTS   sql.NullTime   `json:"updated_ts"`
	UpdatedBy   sql.NullString `json:"updated_by"`
	DeletedTS   sql.NullTime   `json:"deleted_ts"`
	DeletedBy   sql.NullString `json:"deleted_by"`
	Version     int            `json:"version"`
}
