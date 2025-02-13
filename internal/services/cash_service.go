package services

import (
	"context"
	"errors"
	"strconv"

	"github.com/tksasha/balance/internal/apperrors"
)

type CashService struct {
	cashRepository CashRepository
}

func NewCashService(cashRepository CashRepository) *CashService {
	return &CashService{
		cashRepository: cashRepository,
	}
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
