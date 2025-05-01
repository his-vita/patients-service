package model

import (
	"github.com/google/uuid"
	"github.com/his-vita/patients-service/internal/entity"
)

type CreateContact struct {
	PatientID       *uuid.UUID `json:"patient_id"`
	WorkPhoneNumber *string    `json:"work_phone_number"`
	PhoneNumber     *string    `json:"phone_number"`
	Email           *string    `json:"email"`
}

func (c *CreateContact) ToEntity() *entity.Contact {
	return &entity.Contact{
		PatientID:       c.PatientID,
		PhoneNumber:     c.PhoneNumber,
		WorkPhoneNumber: c.WorkPhoneNumber,
		Email:           c.Email,
	}
}

type UpdateContact struct {
	PatientID       *uuid.UUID `json:"patient_id"`
	WorkPhoneNumber *string    `json:"work_phone_number"`
	PhoneNumber     *string    `json:"phone_number"`
	Email           *string    `json:"email"`
}

func (c *UpdateContact) ToEntity() *entity.Contact {
	return &entity.Contact{
		PatientID:       c.PatientID,
		PhoneNumber:     c.PhoneNumber,
		WorkPhoneNumber: c.WorkPhoneNumber,
		Email:           c.Email,
	}
}

type GetContact struct {
	WorkPhoneNumber *string `json:"work_phone_number"`
	PhoneNumber     *string `json:"phone_number"`
	Email           *string `json:"email"`
}
