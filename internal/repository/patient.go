package repository

import (
	"github.com/jackc/pgx/v5"
)

type PatientRepository struct {
	pgCon *pgx.Conn
}

func NewPatientRepository(pgCon *pgx.Conn) *PatientRepository {
	return &PatientRepository{
		pgCon: pgCon,
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
