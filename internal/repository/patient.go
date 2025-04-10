package repository

import (
	"github.com/his-vita/patients-service/internal/database"
)

type PatientRepository struct {
	pgContext *database.PgContext
}

func NewPatientRepository(pgContext *database.PgContext) *PatientRepository {
	return &PatientRepository{
		pgContext: pgContext,
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
