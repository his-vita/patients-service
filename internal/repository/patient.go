package repository

import (
	"fmt"
	"log/slog"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/his-vita/patients-service/internal/database"
	"github.com/his-vita/patients-service/models"
	"github.com/his-vita/patients-service/pkg/sqlutils"
	"github.com/jackc/pgx"
)

type PatientRepository struct {
	log       *slog.Logger
	pgContext *database.PgContext
	sqlFiles  map[string]string
}

func NewPatientRepository(log *slog.Logger, pgContext *database.PgContext, sqlPath string) *PatientRepository {
	filePath := filepath.Join(sqlPath, "patients")
	if err := sqlutils.CheckSQLFilesPath(filePath); err != nil {
		panic(err)
	}

	sqlFiles, err := sqlutils.LoadSQLFiles(filePath)
	if err != nil {
		panic(err)
	}

	return &PatientRepository{
		log:       log,
		pgContext: pgContext,
		sqlFiles:  sqlFiles,
	}
}

func (pr *PatientRepository) GetPatient(id *uuid.UUID) (*models.Patient, error) {
	var patient models.Patient

	ctx, cancel := pr.pgContext.DefaultTimeoutCtx()
	defer cancel()

	query, exists := pr.sqlFiles["get_patient_by_id.sql"]
	if !exists {
		return nil, fmt.Errorf("SQL query insert_patient not found")
	}

	err := pr.pgContext.Pool.QueryRow(ctx, query, id).Scan(
		&patient.Id,
		&patient.FirstName,
		&patient.LastName,
		&patient.MiddleName,
		&patient.BirthDate,
		&patient.PhoneNumber,
		&patient.Email,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("patient with id %s not found", id)
		}
		return nil, fmt.Errorf("error retrieving patient: %w", err)
	}

	return &patient, nil
}

func (pr *PatientRepository) GetAllPatients() {
	panic("impl me!")
}

func (pr *PatientRepository) UpdatePatient() {
	panic("impl me!")
}

func (pr *PatientRepository) CreatePatient(patient *models.Patient) error {
	fmt.Println(pr.sqlFiles)
	query, exists := pr.sqlFiles["insert_patient.sql"]
	if !exists {
		return fmt.Errorf("SQL query insert_patient not found")
	}

	ctx, cancel := pr.pgContext.DefaultTimeoutCtx()
	defer cancel()

	_, err := pr.pgContext.Pool.Exec(ctx, query, patient.FirstName, patient.LastName, patient.MiddleName, patient.BirthDate, patient.PhoneNumber, patient.Email)
	if err != nil {
		return fmt.Errorf("error creating patient: %w", err)
	}

	return nil
}

func (pr *PatientRepository) DeletePatient() {
	panic("impl me!")
}
