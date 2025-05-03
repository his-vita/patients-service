package service

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"github.com/his-vita/patients-service/internal/model"
)

type PatientRepository interface {
	GetPatient(id *uuid.UUID) (*model.Patient, error)
	GetPatients(limit int, offset int) ([]model.Patient, error)
	UpdatePatient(tx context.Context, id *uuid.UUID, patient *model.Patient) error
	CreatePatient(tx context.Context, patient *model.Patient) (*uuid.UUID, error)
	MarkPatientAsDeleted(id *uuid.UUID) error
	UnMarkPatientAsDeleted(id *uuid.UUID) error
}

type PatientTransactionRepository interface {
	CreatePatient(patient *model.Patient) error
}

type PatientService struct {
	log               *slog.Logger
	patientRepository PatientRepository
}

func NewPatientService(log *slog.Logger, r PatientRepository) *PatientService {
	return &PatientService{
		log:               log,
		patientRepository: r,
	}
}

func (ps *PatientService) CreatePatient(tx context.Context, patient *model.Patient) (*uuid.UUID, error) {
	id, err := ps.patientRepository.CreatePatient(tx, patient)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (ps *PatientService) UpdatePatient(tx context.Context, id *uuid.UUID, patient *model.Patient) error {
	err := ps.patientRepository.UpdatePatient(tx, id, patient)
	if err != nil {
		return err
	}

	return nil
}

func (ps *PatientService) GetPatient(id *uuid.UUID) (*model.Patient, error) {
	patient, err := ps.patientRepository.GetPatient(id)
	if err != nil {
		return nil, err
	}

	return patient, nil
}

func (ps *PatientService) GetPatients(limit int, offset int) ([]model.Patient, error) {
	patients, err := ps.patientRepository.GetPatients(limit, offset)
	if err != nil {
		return nil, err
	}

	return patients, nil
}

func (ps *PatientService) MarkPatientAsDeleted(id *uuid.UUID) error {
	err := ps.patientRepository.MarkPatientAsDeleted(id)
	if err != nil {
		return err
	}

	return nil
}

func (ps *PatientService) UnMarkPatientAsDeleted(id *uuid.UUID) error {
	err := ps.patientRepository.UnMarkPatientAsDeleted(id)
	if err != nil {
		return err
	}

	return nil
}
