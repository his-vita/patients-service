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

func PatientDetailsToDTO(entity *entity.Patient) *dto.PatientDetails {
	return &dto.PatientDetails{
		Patient: PatientToDTO(entity),
		Contact: ContactToDTO(&entity.Contact),
	}
}

func PatientDetailsDTOs(entities *[]entity.Patient) *[]dto.PatientDetails {
	patientDTOs := make([]dto.PatientDetails, len(*entities))

	for i, patient := range *entities {
		patientDTO := PatientDetailsToDTO(&patient)
		patientDTOs[i] = *patientDTO
	}

	return &patientDTOs
}
