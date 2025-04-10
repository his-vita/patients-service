package repository

import (
	"path/filepath"

	"github.com/his-vita/patients-service/internal/database"
	"github.com/his-vita/patients-service/pkg/sqlutils"
)

type PatientRepository struct {
	pgContext *database.PgContext
	sqlPath   string
}

func NewPatientRepository(pgContext *database.PgContext, sqlPath string) *PatientRepository {
	filePath := filepath.Join(sqlPath, "patients")
	if err := sqlutils.CheckSQLFilesPath(filePath); err != nil {
		panic(err)
	}

	return &PatientRepository{
		pgContext: pgContext,
		sqlPath:   filePath,
	}
}

func (pr *PatientRepository) GetPatient() {
	panic("impl me!")
}

func (pr *PatientRepository) GetAllPatients() {
	panic("impl me!")
}

func (pr *PatientRepository) UpdatePatient() {
	panic("impl me!")
}

func (pr *PatientRepository) CreatePatient() {
	panic("impl me!")
}

func (pr *PatientRepository) DeletePatient() {
	panic("impl me!")
}
