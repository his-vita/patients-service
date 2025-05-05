package transaction

import (
	"context"
	"fmt"

	"github.com/google/uuid"
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

	if patient.InsuranceOMS != nil {
		patient.InsuranceOMS.PatientID = id

		if err := t.insuranceService.CreateInsurance(tx, patient.InsuranceOMS); err != nil {
			return fmt.Errorf("failed to save insurance OMS: %w", err)
		}
	}

	if patient.InsuranceDMS != nil {
		patient.InsuranceDMS.PatientID = id

		if err := t.insuranceService.CreateInsurance(tx, patient.InsuranceDMS); err != nil {
			return fmt.Errorf("failed to save insurance DMS: %w", err)
		}
	}

	if patient.Document != nil {
		patient.Document.PatientID = id

		if err := t.documentService.CreateDocument(tx, patient.Document); err != nil {
			return fmt.Errorf("failed to save document: %w", err)
		}
	}

	if err := t.txManager.CommitTransaction(tx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (t *Transaction) UpdatePatient(id *uuid.UUID, patient *model.Patient) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tx, err := t.txManager.BeginTransaction(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer t.txManager.RollbackTransaction(tx)

	err = t.patientService.UpdatePatient(tx, id, patient)
	if err != nil {
		return fmt.Errorf("failed to update patient: %w", err)
	}

	if err := t.contactService.UpdateContact(tx, id, &patient.Contact); err != nil {
		return fmt.Errorf("failed to update contact: %w", err)
	}

	if err := t.snilsService.UpdateSnils(tx, id, &patient.Snils); err != nil {
		return fmt.Errorf("failed to update snils: %w", err)
	}

	if err := t.innService.UpdateInn(tx, id, &patient.Inn); err != nil {
		return fmt.Errorf("failed to update inn: %w", err)
	}

	if patient.InsuranceOMS != nil {
		if err := t.insuranceService.UpdateInsurance(tx, patient.InsuranceOMS.ID, patient.InsuranceOMS); err != nil {
			return fmt.Errorf("failed to update insurance: %w", err)
		}
	}

	if patient.InsuranceDMS != nil {
		if err := t.insuranceService.UpdateInsurance(tx, patient.InsuranceDMS.ID, patient.InsuranceDMS); err != nil {
			return fmt.Errorf("failed to update insurance: %w", err)
		}
	}

	if patient.Document != nil {
		if err := t.documentService.UpdateDocument(tx, patient.Document.ID, patient.Document); err != nil {
			return fmt.Errorf("failed to update document: %w", err)
		}
	}

	if err := t.txManager.CommitTransaction(tx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
