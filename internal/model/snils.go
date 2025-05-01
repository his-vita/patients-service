package model

import (
	"github.com/google/uuid"
	"github.com/his-vita/patients-service/internal/entity"
)

type Snils struct {
	PatientID *uuid.UUID `json:"patient_id"`
	Number    *string    `json:"number"`
}

func (s *Snils) ToEntity() *entity.Snils {
	return &entity.Snils{
		PatientID: s.PatientID,
		Number:    s.Number,
	}
}

func (s *Snils) ToModel(snils entity.Snils) *Snils {
	s.Number = snils.Number
	return s
}
