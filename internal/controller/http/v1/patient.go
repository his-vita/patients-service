package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/his-vita/patients-service/internal/model"
)

type PatientService interface {
	GetPatient(id *uuid.UUID) (*model.Patient, error)
	GetPatients(limit int, offset int) ([]model.Patient, error)
	MarkPatientAsDeleted(id *uuid.UUID) error
	UnMarkPatientAsDeleted(id *uuid.UUID) error
}

type Transaction interface {
	CreatePatient(patient *model.Patient) error
	UpdatePatient(id *uuid.UUID, patient *model.Patient) error
}

type PatientController struct {
	patientService PatientService
	transaction    Transaction
}

func NewPatientController(s PatientService, tr Transaction) *PatientController {
	return &PatientController{
		patientService: s,
		transaction:    tr,
	}
}

func (pc *PatientController) CreatePatient(c *gin.Context) {
	var patient model.Patient

	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	if err := pc.transaction.CreatePatient(&patient); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Patient created successfully with transaction"})
}

func (pc *PatientController) UpdatePatient(c *gin.Context) {
	uuid, _ := uuid.Parse(c.Param("id"))

	var patient model.Patient

	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	if err := pc.transaction.UpdatePatient(&uuid, &patient); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Patient updated successfully"})
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

	patients, err := pc.patientService.GetPatients(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, patients)
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
