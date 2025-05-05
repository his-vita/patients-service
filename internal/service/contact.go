package service

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"github.com/his-vita/patients-service/internal/model"
)

type ContactRepository interface {
	CreateContact(tx context.Context, id *uuid.UUID, contact *model.Contact) error
	UpdateContact(tx context.Context, id *uuid.UUID, contact *model.Contact) error
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

func (cs *ContactService) CreateContact(tx context.Context, id *uuid.UUID, contact *model.Contact) error {
	err := cs.contactRepository.CreateContact(tx, id, contact)
	if err != nil {
		return err
	}

	return nil
}

func (cs *ContactService) UpdateContact(tx context.Context, id *uuid.UUID, contact *model.Contact) error {
	err := cs.contactRepository.UpdateContact(tx, id, contact)
	if err != nil {
		return err
	}

	return nil
}
