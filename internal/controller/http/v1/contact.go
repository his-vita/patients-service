package v1

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/his-vita/patients-service/internal/model"
)

type ContactService interface {
	UpdateContact(tx context.Context, updateContact *model.UpdateContact) error
}

type ContactController struct {
	contactService ContactService
}

func NewContactController(s ContactService) *ContactController {
	return &ContactController{
		contactService: s,
	}
}

func (cc *ContactController) UpdateContact(c *gin.Context) {
	var contact model.UpdateContact

	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	if err := cc.contactService.UpdateContact(context.Background(), &contact); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Contact updated successfully"})
}
