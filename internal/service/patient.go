package service

import (
	"github.com/google/uuid"
	"github.com/his-vita/patients-service/internal/entity"
)

type PatientRepository interface {
	GetPatient(id *uuid.UUID) (*entity.Patient, error)
	GetPatients(limit int, offset int) (*[]entity.Patient, error)
	UpdatePatient(patient *entity.Patient) error
	CreatePatient(patient *entity.Patient) error
	MarkPatientAsDeleted(id *uuid.UUID) error
	UnMarkPatientAsDeleted(id *uuid.UUID) error
}

type PatientService struct {
	patientRepository PatientRepository
}

func NewPatientService(r PatientRepository) *PatientService {
	return &PatientService{
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

func (ps *PatientService) GetPatients(limit int, offset int) (*[]entity.Patient, error) {
	patients, err := ps.patientRepository.GetPatients(limit, offset)
	if err != nil {
		return nil, err
	}

	return patients, nil
}

func (ps *PatientService) UpdatePatient(patient *entity.Patient) error {
	err := ps.patientRepository.UpdatePatient(patient)
	if err != nil {
		return err
	}

	return nil
}

func (ps *PatientService) CreatePatient(patient *entity.Patient) error {
	err := ps.patientRepository.CreatePatient(patient)
	if err != nil {
		return err
	}

	return nil
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
