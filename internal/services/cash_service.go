package services

import (
	"context"
	"errors"
	"strconv"

	"github.com/tksasha/balance/internal/apperrors"
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/requests"
	"github.com/tksasha/balance/pkg/validation"
)

type CashService struct {
	cashRepository CashRepository
}

func NewCashService(cashRepository CashRepository) *CashService {
	return &CashService{
		cashRepository: cashRepository,
	}
}

func (s *CashService) Create(ctx context.Context, request requests.CreateCashRequest) error {
	cash := &models.Cash{}

	validate := validation.New()

	cash.Name = validate.Presence("name", request.Name)

	if cash.Name != "" {
		exists, err := s.cashRepository.NameExists(ctx, cash.Name, 0)
		if err != nil {
			return err
		}

		if exists {
			validate.Set("name", AlreadyExists)
		}
	}

	cash.Formula, cash.Sum = validate.Formula("formula", request.Formula)

	cash.Supercategory = validate.Integer("supercategory", request.Supercategory)

	cash.Favorite = validate.Boolean("favorite", request.Favorite)

	if validate.HasErrors() {
		return validate.Errors
	}

	return s.cashRepository.Create(ctx, cash)
}

func (s *CashService) findByID(ctx context.Context, input string) (*models.Cash, error) {
	id, err := strconv.Atoi(input)
	if err != nil {
		return nil, apperrors.ErrResourceNotFound
	}

	cash, err := s.cashRepository.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, apperrors.ErrRecordNotFound) {
			return nil, apperrors.ErrResourceNotFound
		}

		return nil, err
	}

	return cash, nil
}

func (s *CashService) FindByID(ctx context.Context, input string) (*models.Cash, error) {
	return s.findByID(ctx, input)
}

func (s *CashService) Update(ctx context.Context, request requests.CashUpdateRequest) (*models.Cash, error) {
	cash, err := s.findByID(ctx, request.ID)
	if err != nil {
		return nil, err
	}

	validate := validation.New()

	cash.Formula, cash.Sum = validate.Formula("sum", request.Formula)

	cash.Name = validate.Presence("name", request.Name)

	if cash.Name != "" {
		exists, err := s.cashRepository.NameExists(ctx, cash.Name, cash.ID)
		if err != nil {
			return nil, err
		}

		if exists {
			validate.Set("name", AlreadyExists)
		}
	}

	cash.Supercategory = validate.Integer("supercategory", request.Supercategory)

	cash.Favorite = validate.Boolean("favorite", request.Favorite)

	if validate.HasErrors() {
		return cash, validate.Errors
	}

	if err := s.cashRepository.Update(ctx, cash); err != nil {
		if errors.Is(err, apperrors.ErrRecordNotFound) {
			return nil, apperrors.ErrResourceNotFound
		}

		return nil, err
	}

	return cash, nil
}

func (s *CashService) Delete(ctx context.Context, input string) error {
	id, err := strconv.Atoi(input)
	if err != nil {
		return apperrors.ErrResourceNotFound
	}

	if err := s.cashRepository.Delete(ctx, id); err != nil {
		if errors.Is(err, apperrors.ErrRecordNotFound) {
			return apperrors.ErrResourceNotFound
		}

		return err
	}

	return nil
}
