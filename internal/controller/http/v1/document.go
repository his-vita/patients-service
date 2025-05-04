package v1

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/his-vita/patients-service/internal/model"
)

type DocumentService interface {
	CreateDocument(tx context.Context, document *model.Document) error
	UpdateDocument(tx context.Context, id *uuid.UUID, document *model.Document) error
	DeleteDocument(id *uuid.UUID) error
}

type DocumentController struct {
	documentService DocumentService
}

func NewDocumentController(s DocumentService) *DocumentController {
	return &DocumentController{
		documentService: s,
	}
}

func (cr *DocumentController) CreateDocument(c *gin.Context) {
	var document model.Document

	if err := c.ShouldBindJSON(&document); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	if err := cr.documentService.CreateDocument(context.Background(), &document); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Document created successfully"})
}

func (cr *DocumentController) UpdateDocument(c *gin.Context) {
	uuid, _ := uuid.Parse(c.Param("id"))

	var document model.Document

	if err := c.ShouldBindJSON(&document); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	if err := cr.documentService.UpdateDocument(context.Background(), &uuid, &document); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Document updated successfully"})
}

func (cr *DocumentController) DeleteDocument(c *gin.Context) {
	id, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID not found in context"})
		return
	}
	uuid := id.(uuid.UUID)

	if err := cr.documentService.DeleteDocument(&uuid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Document deleted successfully"})
}
