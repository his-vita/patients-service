package model

import (
	"cloud.google.com/go/civil"
	"github.com/google/uuid"
)

type Patient struct {
	ID           uuid.UUID  `json:"id"`
	FirstName    string     `json:"first_name"`
	LastName     string     `json:"last_name"`
	MiddleName   *string    `json:"middle_name"`
	BirthDate    civil.Date `json:"birth_date"`
	Gender       *bool      `json:"gender"`
	Version      int        `json:"version"`
	Contact      Contact    `json:"contact"`
	Snils        Snils      `json:"snils"`
	Inn          Inn        `json:"inn"`
	InsuranceOMS *Insurance `json:"insurance_oms"`
	InsuranceDMS *Insurance `json:"insurance_dms"`
	Document     *Document  `json:"document"`
}

func (p *Patient) Sanitize() {
	if p.InsuranceOMS.ID == nil {
		p.InsuranceOMS = nil
	}

	if p.InsuranceDMS.ID == nil {
		p.InsuranceDMS = nil
	}

	if p.Document.ID == nil {
		p.Document = nil
	}
}
