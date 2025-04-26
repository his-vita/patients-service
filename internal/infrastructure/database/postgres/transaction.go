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
	fmt.Println("rollback1")

	if tx == nil {
		return fmt.Errorf("transaction is already nil, can't rollback")
	}

	err := tx.Rollback(context.Background())
	if err != nil {
		fmt.Printf("error during rollback: %v\n", err)
		return fmt.Errorf("failed to rollback transaction: %w", err)
	}

	fmt.Println("rollback2")
	return nil
}
