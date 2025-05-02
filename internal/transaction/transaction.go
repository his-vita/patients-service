package transaction

import (
	"context"

	"github.com/google/uuid"
	"github.com/his-vita/patients-service/internal/model"
	"github.com/his-vita/patients-service/pkg/database"
)

type PatientService interface {
	CreatePatient(ctx context.Context, patient *model.Patient) (*uuid.UUID, error)
	UpdatePatient(ctx context.Context, patient *model.Patient) error
}

type ContactService interface {
	CreateContact(ctx context.Context, id *uuid.UUID, contact *model.Contact) error
	UpdateContact(ctx context.Context, id *uuid.UUID, contact *model.Contact) error
}

type SnilsService interface {
	CreateSnils(ctx context.Context, id *uuid.UUID, snils *model.Snils) error
	UpdateSnils(ctx context.Context, id *uuid.UUID, snils *model.Snils) error
}

type Transaction struct {
	patientService PatientService
	contactService ContactService
	snilsService   SnilsService
	txManager      database.TransactionManager
}

func NewTransaction(ps PatientService, cs ContactService, ss SnilsService, tx database.TransactionManager) *Transaction {
	return &Transaction{
		patientService: ps,
		contactService: cs,
		snilsService:   ss,
		txManager:      tx,
	}
}
