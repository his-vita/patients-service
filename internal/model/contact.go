package model

import (
	"github.com/his-vita/patients-service/internal/entity"
)

type CreateContact struct {
	WorkPhoneNumber *string `json:"work_phone_number"`
	PhoneNumber     *string `json:"phone_number"`
	Email           *string `json:"email"`
}

func (c *CreateContact) ToEntity() *entity.Contact {
	return &entity.Contact{
		PhoneNumber:     c.PhoneNumber,
		WorkPhoneNumber: c.WorkPhoneNumber,
		Email:           c.Email,
	}
}

type UpdateContact struct {
	WorkPhoneNumber *string `json:"work_phone_number"`
	PhoneNumber     *string `json:"phone_number"`
	Email           *string `json:"email"`
}

func (c *UpdateContact) ToEntity() *entity.Contact {
	return &entity.Contact{
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
