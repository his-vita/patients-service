package v1

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/his-vita/patients-service/internal/model"
)

type InsuranceService interface {
	CreateInsurance(tx context.Context, insurance *model.Insurance) error
	UpdateInsurance(tx context.Context, id *uuid.UUID, insurance *model.Insurance) error
	DeleteInsurance(id *uuid.UUID) error
}

type InsuranceController struct {
	insuranceService InsuranceService
}

func NewInsuranceController(s InsuranceService) *InsuranceController {
	return &InsuranceController{
		insuranceService: s,
	}
}

func (cr *InsuranceController) CreateInsurance(c *gin.Context) {
	var insurance model.Insurance

	if err := c.ShouldBindJSON(&insurance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	if err := cr.insuranceService.CreateInsurance(context.Background(), &insurance); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Insurance created successfully"})
}

func (cr *InsuranceController) UpdateInsurance(c *gin.Context) {
	uuid, _ := uuid.Parse(c.Param("id"))

	var insurance model.Insurance

	if err := c.ShouldBindJSON(&insurance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	if err := cr.insuranceService.UpdateInsurance(context.Background(), &uuid, &insurance); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Insurance updated successfully"})
}

func (cr *InsuranceController) DeleteInsurance(c *gin.Context) {
	id, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID not found in context"})
		return
	}
	uuid := id.(uuid.UUID)

	if err := cr.insuranceService.DeleteInsurance(&uuid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Insurance deleted successfully"})
}
