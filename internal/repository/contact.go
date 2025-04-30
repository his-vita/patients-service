package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/his-vita/patients-service/internal/entity"
	"github.com/his-vita/patients-service/pkg/database/postgres"
	"github.com/his-vita/patients-service/pkg/sqlstore"
	"github.com/jackc/pgx/v5"
)

type ContactRepository struct {
	pgContext *postgres.PgContext
	sqlStore  *sqlstore.SqlStore
}

func NewContactRepository(pgContext *postgres.PgContext, sqlStore *sqlstore.SqlStore) *ContactRepository {
	return &ContactRepository{
		pgContext: pgContext,
		sqlStore:  sqlStore,
	}
}

func (cr *ContactRepository) GetContactsByPatientId(id *uuid.UUID) (*[]entity.Contact, error) {
	query, err := cr.sqlStore.GetQuery("get_contacts_by_patient_id.sql")
	if err != nil {
		return nil, fmt.Errorf("SQL query get_contacts_by_patient_id.sql not found")
	}

	ctx, cancel := cr.pgContext.DefaultTimeoutCtx()
	defer cancel()

	rows, err := cr.pgContext.Pool.Query(ctx, query, id)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	contacts, err := pgx.CollectRows(rows, pgx.RowToStructByName[entity.Contact])
	if err != nil {
		return nil, fmt.Errorf("failed collecting rows: %w", err)
	}

	return &contacts, nil
}

func (cr *ContactRepository) UpdateContact(tx context.Context, contact *entity.Contact) error {
	query, err := cr.sqlStore.GetQuery("update_contact.sql")
	if err != nil {
		return fmt.Errorf("SQL query update_contact.sql not found")
	}

	ctx, cancel := cr.pgContext.DefaultTimeoutCtx()
	defer cancel()

	_, err = cr.pgContext.TxOrDb(tx).Exec(ctx, query,
		contact.ID,
		contact.PhoneNumber,
		contact.WorkPhoneNumber,
		contact.Email)
	if err != nil {
		return fmt.Errorf("error update contact: %w", err)
	}

	return nil
}

func (cr *ContactRepository) CreateContact(tx context.Context, contact *entity.Contact) error {
	query, err := cr.sqlStore.GetQuery("insert_contact.sql")
	if err != nil {
		return fmt.Errorf("SQL query insert_contact.sql not found")
	}

	ctx, cancel := cr.pgContext.DefaultTimeoutCtx()
	defer cancel()

	_, err = cr.pgContext.TxOrDb(tx).Exec(ctx, query,
		contact.PatientID,
		contact.PhoneNumber,
		contact.WorkPhoneNumber,
		contact.Email)
	if err != nil {
		return fmt.Errorf("error creating contact: %w", err)
	}

	return nil
}

func (cr *ContactRepository) DeleteContact(id *uuid.UUID) error {
	query, err := cr.sqlStore.GetQuery("delete_contact.sql")
	if err != nil {
		return fmt.Errorf("SQL query delete_contact.sql not found")
	}

	ctx, cancel := cr.pgContext.DefaultTimeoutCtx()
	defer cancel()

	_, err = cr.pgContext.Pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error deleting contact: %w", err)
	}

	return nil
}
