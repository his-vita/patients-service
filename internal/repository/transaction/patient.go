package transaction

import (
	"fmt"

	"github.com/his-vita/patients-service/internal/entity"
	"github.com/his-vita/patients-service/internal/infrastructure/database"
	"github.com/his-vita/patients-service/internal/infrastructure/sqlstore"
	"github.com/jackc/pgx/v5"
)

type PatientTransactionRepository struct {
	pgContext *database.PgContext
	sqlStore  *sqlstore.SqlStore
}

func NewPatientTransactionRepository(pgContext *database.PgContext, sqlStore *sqlstore.SqlStore) *PatientTransactionRepository {
	return &PatientTransactionRepository{
		pgContext: pgContext,
		sqlStore:  sqlStore,
	}
}

func (r *PatientTransactionRepository) CreatePatient(patient *entity.Patient) error {
	tx, err := r.pgContext.BeginTransaction()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer r.pgContext.RollbackTransaction(tx)

	fmt.Println("run", tx)

	err = r.savePatient(patient, tx)
	if err != nil {
		return fmt.Errorf("failed to save patient: %w", err)
	}

	err = r.saveContact(patient.Contact, tx)
	if err != nil {
		return fmt.Errorf("failed to save patient: %w", err)
	}

	err = r.pgContext.CommitTransaction(tx)
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	fmt.Println("commit", tx)

	return nil
}

func (r *PatientTransactionRepository) savePatient(patient *entity.Patient, tx pgx.Tx) error {
	query, err := r.sqlStore.GetQuery("insert_patient.sql")
	if err != nil {
		return fmt.Errorf("SQL query insert_patient.sql not found")
	}

	ctx, cancel := r.pgContext.DefaultTimeoutCtx()
	defer cancel()

	_, err = tx.Exec(ctx, query,
		patient.FirstName,
		patient.LastName,
		patient.MiddleName,
		patient.BirthDate,
		patient.Gender,
		"admin")
	if err != nil {
		return fmt.Errorf("error creating patient: %w", err)
	}

	fmt.Println("save", tx)

	return nil
}

func (r *PatientTransactionRepository) saveContact(contact *entity.Contact, tx pgx.Tx) error {
	query, err := r.sqlStore.GetQuery("insert_contact.sql")
	if err != nil {
		return fmt.Errorf("SQL query insert_contact.sql not found")
	}

	ctx, cancel := r.pgContext.DefaultTimeoutCtx()
	defer cancel()

	_, err = tx.Exec(ctx, query,
		contact.PatientID,
		contact.PhoneNumber,
		contact.Email,
		contact.Main)
	if err != nil {
		return fmt.Errorf("error creating contact: %w", err)
	}

	return nil
}
