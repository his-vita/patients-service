package mapper

import (
	"github.com/his-vita/patients-service/internal/dto"
	"github.com/his-vita/patients-service/internal/entity"
)

func ContactToEntity(dto *dto.Contact) *entity.Contact {
	return &entity.Contact{
		ID:              dto.ID,
		PatientID:       dto.PatientId,
		PhoneNumber:     dto.PhoneNumber,
		WorkPhoneNumber: dto.WorkPhoneNumber,
		Email:           dto.Email,
	}
}

func ContactToDTO(entity *entity.Contact) *dto.Contact {
	return &dto.Contact{
		ID:          entity.ID,
		PatientId:   entity.PatientID,
		PhoneNumber: entity.PhoneNumber,
		Email:       entity.Email,
	}
}
