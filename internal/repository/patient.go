package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/his-vita/patients-service/internal/model"
	"github.com/his-vita/patients-service/pkg/database/postgres"
	"github.com/his-vita/patients-service/pkg/sqlstore"
	"github.com/jackc/pgx"
)

type PatientRepository struct {
	pgContext *postgres.PgContext
	sqlStore  *sqlstore.SqlStore
}

func NewPatientRepository(pgContext *postgres.PgContext, sqlStore *sqlstore.SqlStore) *PatientRepository {
	return &PatientRepository{
		pgContext: pgContext,
		sqlStore:  sqlStore,
	}
}

func (pr *PatientRepository) CreatePatient(tx context.Context, patient *model.Patient) (*uuid.UUID, error) {
	query, err := pr.sqlStore.GetQuery("insert_patient.sql")
	if err != nil {
		return nil, fmt.Errorf("SQL query insert_patient.sql not found")
	}

	var patientID uuid.UUID

	ctx, cancel := pr.pgContext.DefaultTimeoutCtx()
	defer cancel()

	err = pr.pgContext.TxOrDb(tx).QueryRow(ctx, query,
		patient.FirstName,
		patient.LastName,
		patient.MiddleName,
		patient.BirthDate,
		patient.Gender,
		"admin").Scan(&patientID)
	if err != nil {
		return nil, fmt.Errorf("error creating patient: %w", err)
	}

	return &patientID, nil
}

func (pr *PatientRepository) GetPatient(id *uuid.UUID) (*model.Patient, error) {
	query, err := pr.sqlStore.GetQuery("get_patient_by_id.sql")
	if err != nil {
		return nil, fmt.Errorf("SQL query get_patient_by_id.sql not found")
	}

	ctx, cancel := pr.pgContext.DefaultTimeoutCtx()
	defer cancel()

	patient := model.Patient{
		Insurance: new(model.Insurance),
		Document:  new(model.Document),
	}

	err = pr.pgContext.Pool.QueryRow(ctx, query, id).Scan(
		&patient.ID,
		&patient.FirstName,
		&patient.LastName,
		&patient.MiddleName,
		&patient.BirthDate,
		&patient.Gender,
		&patient.Version,
		&patient.Contact.PhoneNumber,
		&patient.Contact.WorkPhoneNumber,
		&patient.Contact.Email,
		&patient.Snils.Number,
		&patient.Inn.Number,
		&patient.Insurance.ID,
		&patient.Insurance.Number,
		&patient.Insurance.IssueDate,
		&patient.Insurance.ExpiryDate,
		&patient.Insurance.Type,
		&patient.Insurance.InsuranceCompanyID,
		&patient.Document.ID,
		&patient.Document.Series,
		&patient.Document.Number,
		&patient.Document.DepartmentCode,
		&patient.Document.IssueDate,
		&patient.Document.ExpiryDate,
		&patient.Document.Main,
		&patient.Document.DocumentTypeID,
		&patient.Document.DocumentCompanyID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("patient with id %s not found", id)
		}
		return nil, fmt.Errorf("error retrieving patient: %w", err)
	}

	if patient.Insurance.ID == nil {
		patient.Insurance = nil
	}

	if patient.Document.ID == nil {
		patient.Document = nil
	}

	return &patient, nil
}

func (pr *PatientRepository) GetPatients(limit int, offset int) ([]model.Patient, error) {
	query, err := pr.sqlStore.GetQuery("get_patients.sql")
	if err != nil {
		return nil, fmt.Errorf("SQL query get_patients.sql not found")
	}

	ctx, cancel := pr.pgContext.DefaultTimeoutCtx()
	defer cancel()

	rows, err := pr.pgContext.Pool.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var patients []model.Patient

	for rows.Next() {
		patient := model.Patient{
			Insurance: new(model.Insurance),
		}

		err := rows.Scan(
			&patient.ID,
			&patient.FirstName,
			&patient.LastName,
			&patient.MiddleName,
			&patient.BirthDate,
			&patient.Gender,
			&patient.Contact.PhoneNumber,
			&patient.Contact.WorkPhoneNumber,
			&patient.Contact.Email,
			&patient.Snils.Number,
			&patient.Inn.Number,
			&patient.Insurance.Number,
			&patient.Insurance.Type,
			&patient.Insurance.InsuranceCompanyID)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		patients = append(patients, patient)
	}

	return patients, nil
}

func (pr *PatientRepository) UpdatePatient(tx context.Context, id *uuid.UUID, patient *model.Patient) error {
	query, err := pr.sqlStore.GetQuery("update_patient.sql")
	if err != nil {
		return fmt.Errorf("SQL query update_patient.sql not found")
	}

	ctx, cancel := pr.pgContext.DefaultTimeoutCtx()
	defer cancel()

	res, err := pr.pgContext.TxOrDb(tx).Exec(ctx, query,
		id,
		patient.FirstName,
		patient.LastName,
		patient.MiddleName,
		patient.BirthDate,
		patient.Gender,
		"admin",
		patient.Version)
	if err != nil {
		return fmt.Errorf("error update patient: %w", err)
	}

	rowsAffected := res.RowsAffected()

	if rowsAffected == 0 {
		return fmt.Errorf("error update patient. Version in DB is higher")
	}

	return nil
}

func (pr *PatientRepository) MarkPatientAsDeleted(id *uuid.UUID) error {
	query, err := pr.sqlStore.GetQuery("mark_deleted_patient.sql")
	if err != nil {
		return fmt.Errorf("SQL query mark_deleted_patient.sql not found")
	}

	ctx, cancel := pr.pgContext.DefaultTimeoutCtx()
	defer cancel()

	_, err = pr.pgContext.Pool.Exec(ctx, query, id, "admin")
	if err != nil {
		return fmt.Errorf("error mark deleted patient: %w", err)
	}

	return nil
}

func (pr *PatientRepository) UnMarkPatientAsDeleted(id *uuid.UUID) error {
	query, err := pr.sqlStore.GetQuery("unmark_deleted_patient.sql")
	if err != nil {
		return fmt.Errorf("SQL query unmark_deleted_patient.sql not found")
	}

	ctx, cancel := pr.pgContext.DefaultTimeoutCtx()
	defer cancel()

	_, err = pr.pgContext.Pool.Exec(ctx, query, id, "admin")
	if err != nil {
		return fmt.Errorf("error unmark deleted patient: %w", err)
	}

	return nil
}
