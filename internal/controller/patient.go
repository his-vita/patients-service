package controller

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/his-vita/patients-service/internal/service"
	"github.com/his-vita/patients-service/models"
)

type PatientController struct {
	log            *slog.Logger
	patientService *service.PatientService
}

func NewPatientController(log *slog.Logger, s *service.PatientService) *PatientController {
	return &PatientController{
		log:            log,
		patientService: s,
	}
}

func (pc *PatientController) GetPatient(c *gin.Context) {
	rawID, _ := c.Get("id")
	id := rawID.(uuid.UUID)

	patient, err := pc.patientService.GetPatient(&id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, patient)
}

func (pc *PatientController) GetAllPatients(c *gin.Context) {
	pc.patientService.GetAllPatients()
}

func (pc *PatientController) UpdatePatient(c *gin.Context) {
	pc.patientService.UpdatePatient()
}

func (pc *PatientController) CreatePatient(c *gin.Context) {
	var patient models.Patient

	if err := c.ShouldBindJSON(&patient); err != nil {
		pc.log.Error("CreatePatient", "ShouldBindJSON", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	if err := pc.patientService.CreatePatient(&patient); err != nil {
		pc.log.Error("CreatePatient", "PatientService", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Patient created successfully"})
}

func (pc *PatientController) DeletePatient(c *gin.Context) {
	pc.patientService.DeletePatient()
}
