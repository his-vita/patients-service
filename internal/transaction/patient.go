package transaction

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/his-vita/patients-service/internal/dto"
	"github.com/his-vita/patients-service/internal/infrastructure/database"
)

type PatientService interface {
	CreatePatient(patientDTO *dto.Patient) (*uuid.UUID, error)
}

type ContactService interface {
	CreateContact(contactDTO *dto.Contact) error
}

type PatientTransaction struct {
	patientService PatientService
	contactService ContactService
	txManager      database.TransactionManager
}

func NewPatientTransaction(ps PatientService, cs ContactService, tx database.TransactionManager) *PatientTransaction {
	return &PatientTransaction{
		patientService: ps,
		contactService: cs,
		txManager:      tx,
	}
}

func (pt *PatientTransaction) CreatePatientTransaction(patientDTO *dto.PatientDetails) error {
	tx, err := pt.txManager.BeginTransaction()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer pt.txManager.RollbackTransaction(tx)

	id, err := pt.patientService.CreatePatient(patientDTO.Patient)
	if err != nil {
		return fmt.Errorf("failed to save patient: %w", err)
	}

	//patientDTO.Contact.PatientId = id

	if err := pt.contactService.CreateContact(patientDTO.Contact); err != nil {
		return fmt.Errorf("failed to save contact: %w", err)
	}

	if err := pt.txManager.CommitTransaction(tx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	fmt.Println(id)

	return nil
}
