package mapper

import (
	"github.com/his-vita/patients-service/internal/dto"
	"github.com/his-vita/patients-service/internal/entity"
)

type PatientMapper struct{}

func NewPatientMapper() *PatientMapper {
	return &PatientMapper{}
}

func (m *PatientMapper) ToEntity(dto *dto.PatientDTO) *entity.Patient {
	return &entity.Patient{
		ID:         dto.ID,
		FirstName:  dto.FirstName,
		LastName:   dto.LastName,
		MiddleName: dto.MiddleName,
		BirthDate:  dto.BirthDate,
		Gender:     dto.Gender,
	}
}

func (m *PatientMapper) ToDTO(entity *entity.Patient) *dto.PatientDTO {
	return &dto.PatientDTO{
		ID:         entity.ID,
		FirstName:  entity.FirstName,
		LastName:   entity.LastName,
		MiddleName: entity.MiddleName,
		BirthDate:  entity.BirthDate,
		Gender:     entity.Gender,
		Contact: &dto.ContactDTO{
			PhoneNumber: entity.Contact.PhoneNumber,
			Email:       entity.Contact.Email,
		},
	}
}
