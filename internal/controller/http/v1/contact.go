package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/his-vita/patients-service/internal/entity"
)

type ContactService interface {
	GetContactsByPatientId(id *uuid.UUID) (*[]entity.Contact, error)
	UpdateContact(contact *entity.Contact) error
	CreateContact(contact *entity.Contact) error
	DeleteContact(id *uuid.UUID) error
}

type ContactController struct {
	contactService ContactService
}

func NewContactController(s ContactService) *ContactController {
	return &ContactController{
		contactService: s,
	}
}

func (cc *ContactController) GetContactsByPatientId(c *gin.Context) {
	id, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID not found in context"})
		return
	}
	uuid := id.(uuid.UUID)

	contacts, err := cc.contactService.GetContactsByPatientId(&uuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, contacts)
}

func (cc *ContactController) UpdateContact(c *gin.Context) {
	var contact entity.Contact

	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	if err := cc.contactService.UpdateContact(&contact); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Contact updated successfully"})
}

func (cc *ContactController) CreateContact(c *gin.Context) {
	var contact entity.Contact

	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	if err := cc.contactService.CreateContact(&contact); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Contact created successfully"})
}

func (cc *ContactController) DeleteContact(c *gin.Context) {
	id, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID not found in context"})
		return
	}
	uuid := id.(uuid.UUID)

	if err := cc.contactService.DeleteContact(&uuid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Contact created successfully"})
}
