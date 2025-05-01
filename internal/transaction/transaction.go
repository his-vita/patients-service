package transaction

import (
	"context"

	"github.com/google/uuid"
	"github.com/his-vita/patients-service/internal/model"
	"github.com/his-vita/patients-service/pkg/database"
)

type PatientService interface {
	CreatePatient(ctx context.Context, createPatient *model.CreatePatient) (*uuid.UUID, error)
	UpdatePatient(ctx context.Context, updatePatient *model.UpdatePatient) error
}

type ContactService interface {
	CreateContact(ctx context.Context, createContact *model.CreateContact) error
	UpdateContact(ctx context.Context, updateContact *model.UpdateContact) error
}

type Transaction struct {
	patientService PatientService
	contactService ContactService
	txManager      database.TransactionManager
}

func NewTransaction(ps PatientService, cs ContactService, tx database.TransactionManager) *Transaction {
	return &Transaction{
		patientService: ps,
		contactService: cs,
		txManager:      tx,
	}
}
