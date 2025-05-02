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

func (tm *TransactionManager) BeginTransaction(ctx context.Context) (context.Context, error) {
	tx, err := tm.pgContext.Pool.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}

	ctx = context.WithValue(ctx, transactionKey, tx)

	return ctx, nil
}

func (tm *TransactionManager) CommitTransaction(ctx context.Context) error {
	tx, ok := ctx.Value(transactionKey).(pgx.Tx)
	if !ok {
		return fmt.Errorf("transaction not found in context")
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (tm *TransactionManager) RollbackTransaction(ctx context.Context) error {
	tx, ok := ctx.Value(transactionKey).(pgx.Tx)
	if !ok {
		return fmt.Errorf("transaction not found in context")
	}

	if tx == nil {
		return fmt.Errorf("transaction is already nil, can't rollback")
	}

	err := tx.Rollback(ctx)
	if err != nil {
		return fmt.Errorf("failed to rollback transaction: %w", err)
	}

	return nil
}
