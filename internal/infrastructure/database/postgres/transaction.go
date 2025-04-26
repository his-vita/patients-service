package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type TransactionManager struct {
	pgContext *PgContext
}

func NewTransactionManager(pgContext *PgContext) *TransactionManager {
	return &TransactionManager{
		pgContext: pgContext,
	}
}

func (tm *TransactionManager) BeginTransaction() (pgx.Tx, error) {
	tx, err := tm.pgContext.Pool.Begin(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	return tx, nil
}

func (tm *TransactionManager) CommitTransaction(tx pgx.Tx) error {
	if err := tx.Commit(context.Background()); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	return nil
}

func (tm *TransactionManager) RollbackTransaction(tx pgx.Tx) error {
	if err := tx.Rollback(context.Background()); err != nil {
		return fmt.Errorf("failed to rollback transaction: %w", err)
	}
	return nil
}
