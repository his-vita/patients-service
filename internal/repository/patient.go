package repository

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/his-vita/patients-service/internal/entity"
	"github.com/his-vita/patients-service/internal/infrastructure/database"
	"github.com/his-vita/patients-service/internal/infrastructure/sqlstore"
	"github.com/jackc/pgx/v5"
)

type PatientRepository struct {
	pgContext *database.PgContext
	sqlStore  *sqlstore.SqlStore
}

func NewPatientRepository(pgContext *database.PgContext, sqlStore *sqlstore.SqlStore) *PatientRepository {
	return &PatientRepository{
		pgContext: pgContext,
		sqlStore:  sqlStore,
	}
}

func (pr *PatientRepository) GetPatient(id *uuid.UUID) (*entity.Patient, error) {
	query, err := pr.sqlStore.GetQuery("get_patient_by_id.sql")
	if err != nil {
		return nil, fmt.Errorf("SQL query get_patient_by_id.sql not found")
	}

	var patient entity.Patient
	ctx, cancel := pr.pgContext.DefaultTimeoutCtx()
	defer cancel()

	err = pr.pgContext.Pool.QueryRow(ctx, query, id).Scan(
		&patient.ID,
		&patient.FirstName,
		&patient.LastName,
		&patient.MiddleName,
		&patient.BirthDate,
		&patient.Gender,
		&patient.CreatedTS,
		&patient.CreatedBy,
		&patient.UpdatedTS,
		&patient.UpdatedBy,
		&patient.DeletedTS,
		&patient.DeletedBy,
		&patient.Version,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("patient with id %s not found", id)
		}
		return nil, fmt.Errorf("error retrieving patient: %w", err)
	}

	return &patient, nil
}

func (pr *PatientRepository) GetPatients(limit int, offset int) (*[]entity.Patient, error) {
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

	var patients []entity.Patient

	for rows.Next() {
		var patient entity.Patient
		patient.Contact = &entity.Contact{}

		err := rows.Scan(
			&patient.ID,
			&patient.FirstName,
			&patient.LastName,
			&patient.MiddleName,
			&patient.BirthDate,
			&patient.Gender,
			&patient.Contact.PhoneNumber,
			&patient.Contact.Email,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		patients = append(patients, patient)
	}

	return &patients, nil
}

func (pr *PatientRepository) UpdatePatient(patient *entity.Patient) error {
	query, err := pr.sqlStore.GetQuery("update_patient.sql")
	if err != nil {
		return fmt.Errorf("SQL query update_patient.sql not found")
	}

	ctx, cancel := pr.pgContext.DefaultTimeoutCtx()
	defer cancel()

	res, err := pr.pgContext.Pool.Exec(ctx, query,
		patient.ID,
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

func (pr *PatientRepository) CreatePatient(patient *entity.Patient) error {
	query, err := pr.sqlStore.GetQuery("insert_patient.sql")
	if err != nil {
		return fmt.Errorf("SQL query insert_patient.sql not found")
	}

	ctx, cancel := pr.pgContext.DefaultTimeoutCtx()
	defer cancel()

	_, err = pr.pgContext.Pool.Exec(ctx, query,
		patient.FirstName,
		patient.LastName,
		patient.MiddleName,
		patient.BirthDate,
		patient.Gender,
		"admin")
	if err != nil {
		return fmt.Errorf("error creating patient: %w", err)
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
