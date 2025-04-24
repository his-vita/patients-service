package service

import (
	"log/slog"

	"github.com/google/uuid"
	"github.com/his-vita/patients-service/internal/entity"
)

type ContactRepository interface {
	GetContactsByPatientId(id *uuid.UUID) (*[]entity.Contact, error)
	UpdateContact(contact *entity.Contact) error
	CreateContact(contact *entity.Contact) error
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

func (cs *ContactService) UpdateContact(contact *entity.Contact) error {
	err := cs.contactRepository.UpdateContact(contact)
	if err != nil {
		return err
	}

	return nil
}

func (cs *ContactService) CreateContact(contact *entity.Contact) error {
	err := cs.contactRepository.CreateContact(contact)
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
