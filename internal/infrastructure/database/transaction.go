package database

import "github.com/jackc/pgx/v5"

type TransactionManager interface {
	BeginTransaction() (pgx.Tx, error)
	CommitTransaction(tx pgx.Tx) error
	RollbackTransaction(tx pgx.Tx) error
}
