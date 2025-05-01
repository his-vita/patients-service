package service

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/his-vita/patients-service/internal/entity"
	"github.com/his-vita/patients-service/internal/model"
)

type ContactRepository interface {
	CreateContact(tx context.Context, id *uuid.UUID, contact *entity.Contact) error
	UpdateContact(tx context.Context, id *uuid.UUID, contact *entity.Contact) error
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

func (cs *ContactService) CreateContact(tx context.Context, id *uuid.UUID, createContact *model.CreateContact) error {
	contact := createContact.ToEntity()
	if contact == nil {
		return fmt.Errorf("error on contact mapping")
	}

	err := cs.contactRepository.CreateContact(tx, id, contact)
	if err != nil {
		return err
	}

	return nil
}

func (cs *ContactService) UpdateContact(tx context.Context, id *uuid.UUID, updateContact *model.UpdateContact) error {
	contact := updateContact.ToEntity()
	if contact == nil {
		return fmt.Errorf("error on contact mapping")
	}

	err := cs.contactRepository.UpdateContact(tx, id, contact)
	if err != nil {
		return err
	}

	return nil
}
