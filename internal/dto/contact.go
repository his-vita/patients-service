package dto

import "github.com/google/uuid"

type Contact struct {
	ID          uuid.UUID
	PhoneNumber *string
	Email       *string
	Main        bool
}
