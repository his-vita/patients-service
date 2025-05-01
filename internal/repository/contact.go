package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/his-vita/patients-service/internal/entity"
	"github.com/his-vita/patients-service/pkg/database/postgres"
	"github.com/his-vita/patients-service/pkg/sqlstore"
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

func (cr *ContactRepository) CreateContact(tx context.Context, id *uuid.UUID, contact *entity.Contact) error {
	query, err := cr.sqlStore.GetQuery("insert_contact.sql")
	if err != nil {
		return fmt.Errorf("SQL query insert_contact.sql not found")
	}

	ctx, cancel := cr.pgContext.DefaultTimeoutCtx()
	defer cancel()

	_, err = cr.pgContext.TxOrDb(tx).Exec(ctx, query,
		id,
		contact.PhoneNumber,
		contact.WorkPhoneNumber,
		contact.Email)
	if err != nil {
		return fmt.Errorf("error creating contact: %w", err)
	}

	return nil
}

func (cr *ContactRepository) UpdateContact(tx context.Context, id *uuid.UUID, contact *entity.Contact) error {
	query, err := cr.sqlStore.GetQuery("update_contact.sql")
	if err != nil {
		return fmt.Errorf("SQL query update_contact.sql not found")
	}

	ctx, cancel := cr.pgContext.DefaultTimeoutCtx()
	defer cancel()

	_, err = cr.pgContext.TxOrDb(tx).Exec(ctx, query,
		id,
		contact.PhoneNumber,
		contact.WorkPhoneNumber,
		contact.Email)
	if err != nil {
		return fmt.Errorf("error update contact: %w", err)
	}

	return nil
}
