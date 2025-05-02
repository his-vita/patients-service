package service

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"github.com/his-vita/patients-service/internal/model"
)

type InnRepository interface {
	CreateInn(tx context.Context, id *uuid.UUID, inn *model.Inn) error
	UpdateInn(tx context.Context, id *uuid.UUID, inn *model.Inn) error
}

type InnService struct {
	log           *slog.Logger
	innRepository InnRepository
}

func NewInnService(log *slog.Logger, r InnRepository) *InnService {
	return &InnService{
		log:           log,
		innRepository: r,
	}
}

func (cs *InnService) CreateInn(tx context.Context, id *uuid.UUID, inn *model.Inn) error {
	err := cs.innRepository.CreateInn(tx, id, inn)
	if err != nil {
		return err
	}

	return nil
}

func (cs *InnService) UpdateInn(tx context.Context, id *uuid.UUID, inn *model.Inn) error {
	err := cs.innRepository.UpdateInn(tx, id, inn)
	if err != nil {
		return err
	}

	return nil
}
