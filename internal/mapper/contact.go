package mapper

import (
	"github.com/his-vita/patients-service/internal/dto"
	"github.com/his-vita/patients-service/internal/entity"
)

func ContactToEntity(dto *dto.Contact) *entity.Contact {
	return &entity.Contact{
		PhoneNumber: dto.PhoneNumber,
		Email:       dto.Email,
	}
}

func ContactToDTO(entity *entity.Contact) *dto.Contact {
	return &dto.Contact{
		PhoneNumber: entity.PhoneNumber,
		Email:       entity.Email,
	}
}
