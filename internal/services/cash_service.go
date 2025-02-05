package services

import (
	"context"
	"errors"

	internalerrors "github.com/tksasha/balance/internal/errors"
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
	cash := &models.Cash{
		Name:    request.Name,
		Formula: request.Formula,
	}

	validate := validation.New()

	validate.Presence("name", cash.Name)

	if err := s.nameAlreadyExists(ctx, cash.Name, validate); err != nil {
		return err
	}

	cash.Sum = validate.Formula("formula", cash.Formula)

	cash.Supercategory = validate.Integer("supercategory", request.Supercategory)

	cash.Favorite = validate.Boolean("favorite", request.Favorite)

	if validate.HasErrors() {
		return validate.Errors
	}

	return s.cashRepository.Create(ctx, cash)
}

func (s *CashService) nameAlreadyExists(ctx context.Context, name string, validation *validation.Validation) error {
	if name == "" {
		return nil
	}

	_, err := s.cashRepository.FindByName(ctx, name)

	if errors.Is(err, internalerrors.ErrRecordNotFound) {
		return nil
	}

	if err != nil {
		return err
	}

	validation.Set("name", AlreadyExists)

	return nil
}
