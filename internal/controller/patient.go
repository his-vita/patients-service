package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/his-vita/patients-service/internal/service"
)

type PatientController struct {
	patientService *service.PatientService
}

func NewPatientController(s *service.PatientService) *PatientController {
	return &PatientController{
		patientService: s,
	}
}

func (pc *PatientController) GetPatient(c *gin.Context) {
	pc.patientService.GetPatient()
}

func (pc *PatientController) GetAllPatients(c *gin.Context) {
	pc.patientService.GetAllPatients()
}

func (pc *PatientController) UpdatePatient(c *gin.Context) {
	pc.patientService.UpdatePatient()
}

func (pc *PatientController) CreatePatient(c *gin.Context) {
	pc.patientService.CreatePatient()
}

func (pc *PatientController) DeletePatient(c *gin.Context) {
	pc.patientService.DeletePatient()
}
