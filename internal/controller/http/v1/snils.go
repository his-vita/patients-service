package v1

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/his-vita/patients-service/internal/model"
)

type SnilsService interface {
	UpdateSnils(tx context.Context, id *uuid.UUID, updateSnils *model.Snils) error
}

type SnilsController struct {
	snilsService SnilsService
}

func NewSnilsController(s SnilsService) *SnilsController {
	return &SnilsController{
		snilsService: s,
	}
}

func (cc *SnilsController) UpdateSnils(c *gin.Context) {
	id, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID not found in context"})
		return
	}
	uuid := id.(uuid.UUID)

	var snils model.Snils

	if err := c.ShouldBindJSON(&snils); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	if err := cc.snilsService.UpdateSnils(context.Background(), &uuid, &snils); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Snils updated successfully"})
}
