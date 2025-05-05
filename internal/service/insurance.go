package service

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"github.com/his-vita/patients-service/internal/model"
)

type InsuranceRepository interface {
	CreateInsurance(tx context.Context, insurance *model.Insurance) error
	UpdateInsurance(tx context.Context, id *uuid.UUID, insurance *model.Insurance) error
	DeleteInsurance(id *uuid.UUID) error
}

type InsuranceService struct {
	log                 *slog.Logger
	insuranceRepository InsuranceRepository
}

func NewInsuranceService(log *slog.Logger, insuranceRepository InsuranceRepository) *InsuranceService {
	return &InsuranceService{
		log:                 log,
		insuranceRepository: insuranceRepository,
	}
}

func (s *InsuranceService) CreateInsurance(tx context.Context, insurance *model.Insurance) error {
	err := s.insuranceRepository.CreateInsurance(tx, insurance)
	if err != nil {
		return err
	}

	return nil
}

func (s *InsuranceService) UpdateInsurance(tx context.Context, id *uuid.UUID, insurance *model.Insurance) error {
	err := s.insuranceRepository.UpdateInsurance(tx, id, insurance)
	if err != nil {
		return err
	}

	return nil
}

func (s *InsuranceService) DeleteInsurance(id *uuid.UUID) error {
	err := s.insuranceRepository.DeleteInsurance(id)
	if err != nil {
		return err
	}

	return nil
}
