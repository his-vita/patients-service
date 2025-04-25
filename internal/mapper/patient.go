package mapper

import (
	"github.com/his-vita/patients-service/internal/dto"
	"github.com/his-vita/patients-service/internal/entity"
)

func PatientToEntity(dto *dto.Patient) *entity.Patient {
	return &entity.Patient{
		FirstName:  dto.FirstName,
		LastName:   dto.LastName,
		MiddleName: dto.MiddleName,
		BirthDate:  dto.BirthDate,
		Gender:     dto.Gender,
	}
}

func PatientToDTO(entity *entity.Patient) *dto.Patient {
	return &dto.Patient{
		FirstName:  entity.FirstName,
		LastName:   entity.LastName,
		MiddleName: entity.MiddleName,
		BirthDate:  entity.BirthDate,
		Gender:     entity.Gender,
	}
}
