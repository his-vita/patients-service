package v1

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/his-vita/patients-service/internal/model"
)

type SnilsService interface {
	UpdateSnils(tx context.Context, updateSnils *model.Snils) error
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
	var snils model.Snils

	if err := c.ShouldBindJSON(&snils); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	if err := cc.snilsService.UpdateSnils(context.Background(), &snils); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Snils updated successfully"})
}
