package repository

import (
	"context"
	"fmt"

	"github.com/his-vita/patients-service/internal/entity"
	"github.com/his-vita/patients-service/pkg/database/postgres"
	"github.com/his-vita/patients-service/pkg/sqlstore"
)

type SnilsRepository struct {
	pgContext *postgres.PgContext
	sqlStore  *sqlstore.SqlStore
}

func NewSnilsRepository(pgContext *postgres.PgContext, sqlStore *sqlstore.SqlStore) *SnilsRepository {
	return &SnilsRepository{
		pgContext: pgContext,
		sqlStore:  sqlStore,
	}
}

func (cr *SnilsRepository) CreateSnils(tx context.Context, snils *entity.Snils) error {
	query, err := cr.sqlStore.GetQuery("insert_snils.sql")
	if err != nil {
		return fmt.Errorf("SQL query insert_snils.sql not found")
	}

	ctx, cancel := cr.pgContext.DefaultTimeoutCtx()
	defer cancel()

	_, err = cr.pgContext.TxOrDb(tx).Exec(ctx, query,
		snils.PatientID,
		snils.Number)
	if err != nil {
		return fmt.Errorf("error creating snils: %w", err)
	}

	return nil
}

func (cr *SnilsRepository) UpdateSnils(tx context.Context, snils *entity.Snils) error {
	query, err := cr.sqlStore.GetQuery("update_snils.sql")
	if err != nil {
		return fmt.Errorf("SQL query update_snils.sql not found")
	}

	ctx, cancel := cr.pgContext.DefaultTimeoutCtx()
	defer cancel()

	_, err = cr.pgContext.TxOrDb(tx).Exec(ctx, query,
		snils.PatientID,
		snils.Number)
	if err != nil {
		return fmt.Errorf("error update snils: %w", err)
	}

	return nil
}
