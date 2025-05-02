package transaction

import (
	"context"
	"fmt"

	"github.com/his-vita/patients-service/internal/model"
)

func (t *Transaction) CreatePatient(createPatient *model.Patient) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tx, err := t.txManager.BeginTransaction(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer t.txManager.RollbackTransaction(tx)

	id, err := t.patientService.CreatePatient(tx, createPatient)
	if err != nil {
		return fmt.Errorf("failed to save patient: %w", err)
	}

	if err := t.contactService.CreateContact(tx, id, &createPatient.Contact); err != nil {
		return fmt.Errorf("failed to save contact: %w", err)
	}

	if err := t.snilsService.CreateSnils(tx, id, &createPatient.Snils); err != nil {
		return fmt.Errorf("failed to save snils: %w", err)
	}

	if err := t.txManager.CommitTransaction(tx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (t *Transaction) UpdatePatient(updatePatient *model.Patient) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	fmt.Println(updatePatient)

	tx, err := t.txManager.BeginTransaction(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer t.txManager.RollbackTransaction(tx)

	err = t.patientService.UpdatePatient(tx, updatePatient)
	if err != nil {
		return fmt.Errorf("failed to update patient: %w", err)
	}

	if err := t.contactService.UpdateContact(tx, &updatePatient.ID, &updatePatient.Contact); err != nil {
		return fmt.Errorf("failed to update contact: %w", err)
	}

	if err := t.snilsService.UpdateSnils(tx, &updatePatient.ID, &updatePatient.Snils); err != nil {
		return fmt.Errorf("failed to update snils: %w", err)
	}

	if err := t.txManager.CommitTransaction(tx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
