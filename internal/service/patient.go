package service

import (
	"github.com/google/uuid"
	"github.com/his-vita/patients-service/internal/repository"
	"github.com/his-vita/patients-service/models"
)

type PatientService struct {
	patientRepository *repository.PatientRepository
}

func NewPatientService(r *repository.PatientRepository) *PatientService {
	return &PatientService{
		patientRepository: r,
	}
}

func (ps *PatientService) GetPatient(id *uuid.UUID) (*models.Patient, error) {
	patient, err := ps.patientRepository.GetPatient(id)
	if err != nil {
		return nil, err
	}

	return patient, nil
}

func (ps *PatientService) GetAllPatients(limit int, offset int) ([]models.Patient, error) {
	patients, err := ps.patientRepository.GetPatients(limit, offset)
	if err != nil {
		return nil, err
	}

	return patients, nil
}

func (ps *PatientService) UpdatePatient(patient *models.Patient) error {
	err := ps.patientRepository.UpdatePatient(patient)
	if err != nil {
		return err
	}

	return nil
}

func (ps *PatientService) CreatePatient(patient *models.Patient) error {
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
