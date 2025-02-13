package services

import (
	"context"
	"errors"
	"strconv"

	"github.com/tksasha/balance/internal/apperrors"
	"github.com/tksasha/balance/internal/models"
)

type CashService struct {
	cashRepository CashRepository
}

func NewCashService(cashRepository CashRepository) *CashService {
	return &CashService{
		cashRepository: cashRepository,
	}
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
