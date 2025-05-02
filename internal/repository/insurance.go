package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/his-vita/patients-service/internal/model"
	"github.com/his-vita/patients-service/pkg/database/postgres"
	"github.com/his-vita/patients-service/pkg/sqlstore"
)

type InsuranceRepository struct {
	pgContext *postgres.PgContext
	sqlStore  *sqlstore.SqlStore
}

func NewInsuranceRepository(pgContext *postgres.PgContext, sqlStore *sqlstore.SqlStore) *InsuranceRepository {
	return &InsuranceRepository{
		pgContext: pgContext,
		sqlStore:  sqlStore,
	}
}

func (r *InsuranceRepository) CreateInsurance(tx context.Context, insurance *model.Insurance) error {
	query, err := r.sqlStore.GetQuery("insert_insurance_policies.sql")
	if err != nil {
		return fmt.Errorf("SQL query insert_insurance_policies.sql not found")
	}

	ctx, cancel := r.pgContext.DefaultTimeoutCtx()
	defer cancel()

	_, err = r.pgContext.TxOrDb(tx).Exec(ctx, query,
		insurance.PatientID,
		insurance.Number,
		insurance.IssueDate,
		insurance.ExpiryDate,
		insurance.Type,
		insurance.InsuranceCompanyID)
	if err != nil {
		return fmt.Errorf("error creating insurance policy: %w", err)
	}

	return nil
}

func (r *InsuranceRepository) UpdateInsurance(tx context.Context, id *uuid.UUID, insurance *model.Insurance) error {
	query, err := r.sqlStore.GetQuery("update_insurance_policies.sql")
	if err != nil {
		return fmt.Errorf("SQL query update_insurance_policies.sql not found")
	}

	ctx, cancel := r.pgContext.DefaultTimeoutCtx()
	defer cancel()

	_, err = r.pgContext.TxOrDb(tx).Exec(ctx, query,
		id,
		insurance.Number,
		insurance.IssueDate,
		insurance.ExpiryDate,
		insurance.Type,
		insurance.InsuranceCompanyID)
	if err != nil {
		return fmt.Errorf("error update insurance policy: %w", err)
	}

	return nil
}

func (r *InsuranceRepository) DeleteInsurance(id *uuid.UUID) error {
	query, err := r.sqlStore.GetQuery("delete_insurance_policies.sql")
	if err != nil {
		return fmt.Errorf("SQL query delete_insurance_policies.sql not found")
	}

	ctx, cancel := r.pgContext.DefaultTimeoutCtx()
	defer cancel()

	_, err = r.pgContext.Pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error delete insurance policy")
	}

	return nil
}
