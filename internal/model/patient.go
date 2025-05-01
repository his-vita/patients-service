package model

import (
	"cloud.google.com/go/civil"
	"github.com/google/uuid"
	"github.com/his-vita/patients-service/internal/entity"
)

type CreatePatient struct {
	FirstName  string        `json:"first_name"`
	LastName   string        `json:"last_name"`
	MiddleName *string       `json:"middle_name"`
	BirthDate  civil.Date    `json:"birth_date"`
	Gender     *bool         `json:"gender"`
	Contact    CreateContact `json:"contact"`
}

func (p *CreatePatient) ToEntity() *entity.Patient {
	return &entity.Patient{
		FirstName:  p.FirstName,
		LastName:   p.LastName,
		MiddleName: p.MiddleName,
		BirthDate:  p.BirthDate,
		Gender:     p.Gender,
	}
}

type UpdatePatient struct {
	ID         uuid.UUID     `json:"id"`
	FirstName  string        `json:"first_name"`
	LastName   string        `json:"last_name"`
	MiddleName *string       `json:"middle_name"`
	BirthDate  civil.Date    `json:"birth_date"`
	Gender     *bool         `json:"gender"`
	Version    int           `json:"version"`
	Contact    UpdateContact `json:"contact"`
}

func (p *UpdatePatient) ToEntity() *entity.Patient {
	return &entity.Patient{
		ID:         p.ID,
		FirstName:  p.FirstName,
		LastName:   p.LastName,
		MiddleName: p.MiddleName,
		BirthDate:  p.BirthDate,
		Version:    p.Version,
		Gender:     p.Gender,
	}
}

type GetPatient struct {
	ID         uuid.UUID  `json:"id"`
	FirstName  string     `json:"first_name"`
	LastName   string     `json:"last_name"`
	MiddleName *string    `json:"middle_name"`
	BirthDate  civil.Date `json:"birth_date"`
	Gender     *bool      `json:"gender"`
	Version    int        `json:"version,omitempty"`
	Contact    GetContact `json:"contact"`
}

func (p *GetPatient) ToModel(patient *entity.Patient) *GetPatient {
	return &GetPatient{
		ID:         patient.ID,
		FirstName:  patient.FirstName,
		LastName:   patient.LastName,
		MiddleName: patient.MiddleName,
		BirthDate:  patient.BirthDate,
		Gender:     patient.Gender,
		Version:    patient.Version,
		Contact: GetContact{
			ID:              patient.Contact.ID,
			WorkPhoneNumber: patient.Contact.WorkPhoneNumber,
			PhoneNumber:     patient.Contact.PhoneNumber,
			Email:           patient.Contact.Email,
		},
	}
}

func (p *GetPatient) ToModelList(patients []entity.Patient) []GetPatient {
	patientModels := make([]GetPatient, len(patients))

	for i, patient := range patients {
		patientModels[i] = *p.ToModel(&patient)
	}

	return patientModels
}
