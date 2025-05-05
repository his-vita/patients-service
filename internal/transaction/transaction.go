package transaction

import (
	"context"

	"github.com/google/uuid"
	"github.com/his-vita/patients-service/internal/model"
	"github.com/his-vita/patients-service/pkg/database"
)

type PatientService interface {
	CreatePatient(ctx context.Context, patient *model.Patient) (*uuid.UUID, error)
	UpdatePatient(ctx context.Context, id *uuid.UUID, patient *model.Patient) error
}

type ContactService interface {
	CreateContact(ctx context.Context, id *uuid.UUID, contact *model.Contact) error
	UpdateContact(ctx context.Context, id *uuid.UUID, contact *model.Contact) error
}

type SnilsService interface {
	CreateSnils(ctx context.Context, id *uuid.UUID, snils *model.Snils) error
	UpdateSnils(ctx context.Context, id *uuid.UUID, snils *model.Snils) error
}

type InnService interface {
	CreateInn(ctx context.Context, id *uuid.UUID, inn *model.Inn) error
	UpdateInn(ctx context.Context, id *uuid.UUID, inn *model.Inn) error
}

type InsuranceService interface {
	CreateInsurance(ctx context.Context, insurance *model.Insurance) error
	UpdateInsurance(ctx context.Context, id *uuid.UUID, insurance *model.Insurance) error
}

type DocumentService interface {
	CreateDocument(ctx context.Context, document *model.Document) error
	UpdateDocument(ctx context.Context, id *uuid.UUID, document *model.Document) error
}

type Transaction struct {
	patientService   PatientService
	contactService   ContactService
	snilsService     SnilsService
	innService       InnService
	insuranceService InsuranceService
	documentService  DocumentService
	txManager        database.TransactionManager
}

func NewTransaction(patientService PatientService,
	ContactService ContactService, SnilsService SnilsService,
	InnService InnService, insuranceService InsuranceService,
	documentService DocumentService,
	tx database.TransactionManager) *Transaction {
	return &Transaction{
		patientService:   patientService,
		contactService:   ContactService,
		snilsService:     SnilsService,
		innService:       InnService,
		insuranceService: insuranceService,
		documentService:  documentService,
		txManager:        tx,
	}
}
