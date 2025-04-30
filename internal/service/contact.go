package service

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/his-vita/patients-service/internal/dto"
	"github.com/his-vita/patients-service/internal/entity"
	"github.com/his-vita/patients-service/internal/mapper"
)

type ContactRepository interface {
	GetContactsByPatientId(id *uuid.UUID) (*[]entity.Contact, error)
	UpdateContact(tx context.Context, contact *entity.Contact) error
	CreateContact(tx context.Context, contact *entity.Contact) error
	DeleteContact(id *uuid.UUID) error
}

type ContactService struct {
	log               *slog.Logger
	contactRepository ContactRepository
}

func NewContactService(log *slog.Logger, r ContactRepository) *ContactService {
	return &ContactService{
		log:               log,
		contactRepository: r,
	}
}

func (cs *ContactService) GetContactsByPatientId(id *uuid.UUID) (*[]entity.Contact, error) {
	contacts, err := cs.contactRepository.GetContactsByPatientId(id)
	if err != nil {
		return nil, err
	}

	return contacts, nil
}

func (cs *ContactService) UpdateContact(tx context.Context, contact *entity.Contact) error {
	err := cs.contactRepository.UpdateContact(tx, contact)
	if err != nil {
		return err
	}

	return nil
}

func (cs *ContactService) CreateContact(tx context.Context, contactDTO *dto.Contact) error {
	contact := mapper.ContactToEntity(contactDTO)
	if contact == nil {
		return fmt.Errorf("error on contact mapping")
	}

	err := cs.contactRepository.CreateContact(tx, contact)
	if err != nil {
		return err
	}

	return nil
}

func (cs *ContactService) DeleteContact(id *uuid.UUID) error {
	err := cs.contactRepository.DeleteContact(id)
	if err != nil {
		return err
	}

	return nil
}
