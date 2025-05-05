package service

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"github.com/his-vita/patients-service/internal/model"
)

type DocumentRepository interface {
	CreateDocument(tx context.Context, document *model.Document) error
	UpdateDocument(tx context.Context, id *uuid.UUID, document *model.Document) error
	DeleteDocument(id *uuid.UUID) error
}

type DocumentService struct {
	log                *slog.Logger
	documentRepository DocumentRepository
}

func NewDocumentService(log *slog.Logger, documentRepository DocumentRepository) *DocumentService {
	return &DocumentService{
		log:                log,
		documentRepository: documentRepository,
	}
}

func (s *DocumentService) CreateDocument(tx context.Context, document *model.Document) error {
	err := s.documentRepository.CreateDocument(tx, document)
	if err != nil {
		return err
	}

	return nil
}

func (s *DocumentService) UpdateDocument(tx context.Context, id *uuid.UUID, document *model.Document) error {
	err := s.documentRepository.UpdateDocument(tx, id, document)
	if err != nil {
		return err
	}

	return nil
}

func (s *DocumentService) DeleteDocument(id *uuid.UUID) error {
	err := s.documentRepository.DeleteDocument(id)
	if err != nil {
		return err
	}

	return nil
}
