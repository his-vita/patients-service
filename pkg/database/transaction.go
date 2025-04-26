package database

import (
	"context"
)

type TransactionManager interface {
	BeginTransaction(ctx context.Context) (context.Context, error)
	CommitTransaction(ctx context.Context) error
	RollbackTransaction(ctx context.Context) error
	// BeginTransaction() (pgx.Tx, error)
	// CommitTransaction(tx pgx.Tx) error
	// RollbackTransaction(tx pgx.Tx) error
}
