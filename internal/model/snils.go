package model

import (
	"github.com/his-vita/patients-service/internal/entity"
)

type Snils struct {
	Number *string `json:"number"`
}

func (s *Snils) ToEntity() *entity.Snils {
	return &entity.Snils{
		Number: s.Number,
	}
}

func (s *Snils) ToModel(snils entity.Snils) *Snils {
	s.Number = snils.Number
	return s
}
