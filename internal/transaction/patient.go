package transaction

import (
	"context"
	"fmt"

	"github.com/his-vita/patients-service/internal/model"
)

func (t *Transaction) CreatePatient(patient *model.Patient) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tx, err := t.txManager.BeginTransaction(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer t.txManager.RollbackTransaction(tx)

	id, err := t.patientService.CreatePatient(tx, patient)
	if err != nil {
		return fmt.Errorf("failed to save patient: %w", err)
	}

	if err := t.contactService.CreateContact(tx, id, &patient.Contact); err != nil {
		return fmt.Errorf("failed to save contact: %w", err)
	}

	if err := t.snilsService.CreateSnils(tx, id, &patient.Snils); err != nil {
		return fmt.Errorf("failed to save snils: %w", err)
	}

	if err := t.innService.CreateInn(tx, id, &patient.Inn); err != nil {
		return fmt.Errorf("failed to save inn: %w", err)
	}

	if err := t.txManager.CommitTransaction(tx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (t *Transaction) UpdatePatient(patient *model.Patient) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tx, err := t.txManager.BeginTransaction(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer t.txManager.RollbackTransaction(tx)

	err = t.patientService.UpdatePatient(tx, patient)
	if err != nil {
		return fmt.Errorf("failed to update patient: %w", err)
	}

	if err := t.contactService.UpdateContact(tx, &patient.ID, &patient.Contact); err != nil {
		return fmt.Errorf("failed to update contact: %w", err)
	}

	if err := t.snilsService.UpdateSnils(tx, &patient.ID, &patient.Snils); err != nil {
		return fmt.Errorf("failed to update snils: %w", err)
	}

	if err := t.innService.UpdateInn(tx, &patient.ID, &patient.Inn); err != nil {
		return fmt.Errorf("failed to update snils: %w", err)
	}

	if err := t.txManager.CommitTransaction(tx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
