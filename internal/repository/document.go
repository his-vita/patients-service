package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/his-vita/patients-service/internal/model"
	"github.com/his-vita/patients-service/pkg/database/postgres"
	"github.com/his-vita/patients-service/pkg/sqlstore"
)

type DocumentRepository struct {
	pgContext *postgres.PgContext
	sqlStore  *sqlstore.SqlStore
}

func NewDocumentRepository(pgContext *postgres.PgContext, sqlStore *sqlstore.SqlStore) *DocumentRepository {
	return &DocumentRepository{
		pgContext: pgContext,
		sqlStore:  sqlStore,
	}
}

func (r *DocumentRepository) CreateDocument(tx context.Context, document *model.Document) error {
	query, err := r.sqlStore.GetQuery("insert_document.sql")
	if err != nil {
		return fmt.Errorf("SQL query insert_document.sql not found")
	}

	ctx, cancel := r.pgContext.DefaultTimeoutCtx()
	defer cancel()

	_, err = r.pgContext.TxOrDb(tx).Exec(ctx, query,
		document.PatientID,
		document.Series,
		document.Number,
		document.DepartmentCode,
		document.IssueDate,
		document.ExpiryDate,
		document.Main,
		document.DocumentTypeID,
		document.DocumentCompanyID)
	if err != nil {
		return fmt.Errorf("error creating document: %w", err)
	}

	return nil
}

func (r *DocumentRepository) UpdateDocument(tx context.Context, id *uuid.UUID, document *model.Document) error {
	query, err := r.sqlStore.GetQuery("update_document.sql")
	if err != nil {
		return fmt.Errorf("SQL query update_document.sql not found")
	}

	ctx, cancel := r.pgContext.DefaultTimeoutCtx()
	defer cancel()

	_, err = r.pgContext.TxOrDb(tx).Exec(ctx, query,
		id,
		document.Series,
		document.Number,
		document.DepartmentCode,
		document.IssueDate,
		document.ExpiryDate,
		document.Main,
		document.DocumentTypeID,
		document.DocumentCompanyID)
	if err != nil {
		return fmt.Errorf("error update document policy: %w", err)
	}

	return nil
}

func (r *DocumentRepository) DeleteDocument(id *uuid.UUID) error {
	query, err := r.sqlStore.GetQuery("delete_document.sql")
	if err != nil {
		return fmt.Errorf("SQL query delete_document.sql not found")
	}

	ctx, cancel := r.pgContext.DefaultTimeoutCtx()
	defer cancel()

	_, err = r.pgContext.Pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error delete document")
	}

	return nil
}
