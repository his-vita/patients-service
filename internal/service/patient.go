package service

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"github.com/his-vita/patients-service/internal/dto"
	"github.com/his-vita/patients-service/internal/entity"
	"github.com/his-vita/patients-service/internal/mapper"
)

type PatientRepository interface {
	GetPatient(id *uuid.UUID) (*entity.Patient, error)
	GetPatients(limit int, offset int) (*[]entity.Patient, error)
	UpdatePatient(tx context.Context, patient *entity.Patient) error
	CreatePatient(tx context.Context, patient *entity.Patient) (*uuid.UUID, error)
	MarkPatientAsDeleted(id *uuid.UUID) error
	UnMarkPatientAsDeleted(id *uuid.UUID) error
}

type PatientTransactionRepository interface {
	CreatePatient(patient *entity.Patient) error
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

func (ps *PatientService) GetPatient(id *uuid.UUID) (*entity.Patient, error) {
	patient, err := ps.patientRepository.GetPatient(id)
	if err != nil {
		return nil, err
	}

	return patient, nil
}

func (ps *PatientService) GetPatients(limit int, offset int) (*[]dto.PatientDetails, error) {
	patients, err := ps.patientRepository.GetPatients(limit, offset)
	if err != nil {
		return nil, err
	}

	patientDTOs := mapper.PatientDetailsDTOs(patients)

	return patientDTOs, nil
}

func (ps *PatientService) UpdatePatient(tx context.Context, patient *entity.Patient) error {
	err := ps.patientRepository.UpdatePatient(tx, patient)
	if err != nil {
		return err
	}

	return nil
}

func (ps *PatientService) CreatePatient(tx context.Context, patientDTO *dto.Patient) (*uuid.UUID, error) {
	patient := mapper.PatientToEntity(patientDTO)

	id, err := ps.patientRepository.CreatePatient(tx, patient)
	if err != nil {
		return nil, err
	}

	return id, nil
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
