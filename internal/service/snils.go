package service

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/his-vita/patients-service/internal/entity"
	"github.com/his-vita/patients-service/internal/model"
)

type SnilsRepository interface {
	CreateSnils(tx context.Context, id *uuid.UUID, snils *entity.Snils) error
	UpdateSnils(tx context.Context, id *uuid.UUID, snils *entity.Snils) error
}

type SnilsService struct {
	log             *slog.Logger
	snilsRepository SnilsRepository
}

func NewSnilsService(log *slog.Logger, r SnilsRepository) *SnilsService {
	return &SnilsService{
		log:             log,
		snilsRepository: r,
	}
}

func (cs *SnilsService) CreateSnils(tx context.Context, id *uuid.UUID, createSnils *model.Snils) error {
	snils := createSnils.ToEntity()
	if snils == nil {
		return fmt.Errorf("error on snils mapping")
	}

	err := cs.snilsRepository.CreateSnils(tx, id, snils)
	if err != nil {
		return err
	}

	return nil
}

func (cs *SnilsService) UpdateSnils(tx context.Context, id *uuid.UUID, updateSnils *model.Snils) error {
	snils := updateSnils.ToEntity()
	if snils == nil {
		return fmt.Errorf("error on snils mapping")
	}

	err := cs.snilsRepository.UpdateSnils(tx, id, snils)
	if err != nil {
		return err
	}

	return nil
}
