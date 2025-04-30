package transaction

import (
	"context"

	"github.com/google/uuid"
	"github.com/his-vita/patients-service/internal/dto"
	"github.com/his-vita/patients-service/pkg/database"
)

type PatientService interface {
	CreatePatient(ctx context.Context, patientDTO *dto.Patient) (*uuid.UUID, error)
}

type ContactService interface {
	CreateContact(ctx context.Context, contactDTO *dto.Contact) error
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
