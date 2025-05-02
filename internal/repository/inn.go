package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/his-vita/patients-service/internal/model"
	"github.com/his-vita/patients-service/pkg/database/postgres"
	"github.com/his-vita/patients-service/pkg/sqlstore"
)

type InnRepository struct {
	pgContext *postgres.PgContext
	sqlStore  *sqlstore.SqlStore
}

func NewInnRepository(pgContext *postgres.PgContext, sqlStore *sqlstore.SqlStore) *InnRepository {
	return &InnRepository{
		pgContext: pgContext,
		sqlStore:  sqlStore,
	}
}

func (cr *InnRepository) CreateInn(tx context.Context, id *uuid.UUID, inn *model.Inn) error {
	query, err := cr.sqlStore.GetQuery("insert_inn.sql")
	if err != nil {
		return fmt.Errorf("SQL query insert_inn.sql not found")
	}

	ctx, cancel := cr.pgContext.DefaultTimeoutCtx()
	defer cancel()

	_, err = cr.pgContext.TxOrDb(tx).Exec(ctx, query,
		id,
		inn.Number)
	if err != nil {
		return fmt.Errorf("error creating inn: %w", err)
	}

	return nil
}

func (cr *InnRepository) UpdateInn(tx context.Context, id *uuid.UUID, inn *model.Inn) error {
	query, err := cr.sqlStore.GetQuery("update_inn.sql")
	if err != nil {
		return fmt.Errorf("SQL query update_inn.sql not found")
	}

	ctx, cancel := cr.pgContext.DefaultTimeoutCtx()
	defer cancel()

	_, err = cr.pgContext.TxOrDb(tx).Exec(ctx, query,
		id,
		inn.Number)
	if err != nil {
		return fmt.Errorf("error update inn: %w", err)
	}

	return nil
}
