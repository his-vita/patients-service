package controller

import (
	"log/slog"
	"net/http"
	"strconv"

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
	id, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID not found in context"})
		return
	}
	uuid := id.(uuid.UUID)

	patient, err := pc.patientService.GetPatient(&uuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, patient)
}

func (pc *PatientController) GetPatients(c *gin.Context) {
	limitStr := c.Param("limit")
	offsetStr := c.Param("offset")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset"})
		return
	}

	patients, err := pc.patientService.GetAllPatients(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, patients)
}

func (pc *PatientController) UpdatePatient(c *gin.Context) {
	var patient models.Patient

	if err := c.ShouldBindJSON(&patient); err != nil {
		pc.log.Error("UpdatePatient", "ShouldBindJSON", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	if err := pc.patientService.UpdatePatient(&patient); err != nil {
		pc.log.Error("UpdatePatient", "PatientService", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Patient updated successfully"})
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

func (pc *PatientController) MarkPatientAsDeleted(c *gin.Context) {
	id, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID not found in context"})
		return
	}
	uuid := id.(uuid.UUID)

	err := pc.patientService.MarkPatientAsDeleted(&uuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (pc *PatientController) UnMarkPatientAsDeleted(c *gin.Context) {
	id, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID not found in context"})
		return
	}
	uuid := id.(uuid.UUID)

	err := pc.patientService.UnMarkPatientAsDeleted(&uuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}
