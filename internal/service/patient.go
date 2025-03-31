package service

import (
	"github.com/his-vita/patients-service/internal/repository"
)

type PatientService struct {
	patientRepository *repository.PatientRepository
}

func NewPatientService(r *repository.PatientRepository) *PatientService {
	return &PatientService{
		patientRepository: r,
	}
}

func (ps *PatientService) GetPatient() {
	ps.patientRepository.GetPatient()
}

func (ps *PatientService) GetAllPatients() {
	ps.patientRepository.GetAllPatients()
}

func (ps *PatientService) UpdatePatient() {
	ps.patientRepository.UpdatePatient()
}

func (ps *PatientService) CreatePatient() {
	ps.patientRepository.CreatePatient()
}

func (ps *PatientService) DeletePatient() {
	ps.patientRepository.DeletePatient()
}
