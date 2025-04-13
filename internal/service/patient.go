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

func (ps *PatientService) GetAllPatients() {
	ps.patientRepository.GetAllPatients()
}

func (ps *PatientService) UpdatePatient() {
	ps.patientRepository.UpdatePatient()
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
