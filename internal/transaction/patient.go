package transaction

import (
	"context"
	"fmt"

	"github.com/his-vita/patients-service/internal/dto"
)

func (ts *Transaction) CreatePatient(patientDTO *dto.PatientDetails) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tx, err := ts.txManager.BeginTransaction(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer ts.txManager.RollbackTransaction(tx)

	id, err := ts.patientService.CreatePatient(tx, patientDTO.Patient)
	if err != nil {
		return fmt.Errorf("failed to save patient: %w", err)
	}

	patientDTO.Contact.PatientId = id

	if err := ts.contactService.CreateContact(tx, patientDTO.Contact); err != nil {
		return fmt.Errorf("failed to save contact: %w", err)
	}

	if err := ts.txManager.CommitTransaction(tx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
