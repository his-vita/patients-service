package v1

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/his-vita/patients-service/internal/model"
)

type InnService interface {
	UpdateInn(tx context.Context, id *uuid.UUID, updateInn *model.Inn) error
}

type InnController struct {
	innService InnService
}

func NewInnController(s InnService) *InnController {
	return &InnController{
		innService: s,
	}
}

func (cc *InnController) UpdateInn(c *gin.Context) {
	id, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID not found in context"})
		return
	}
	uuid := id.(uuid.UUID)

	var inn model.Inn

	if err := c.ShouldBindJSON(&inn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	if err := cc.innService.UpdateInn(context.Background(), &uuid, &inn); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inn updated successfully"})
}
