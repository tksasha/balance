package service

import (
	"context"

	"github.com/tksasha/balance/internal/app/balance"
)

func (s *Service) Balance(ctx context.Context) (*balance.Balance, error) {
	residual, err := s.residual(ctx)
	if err != nil {
		return nil, err
	}

	cashes, err := s.repository.Cashes(ctx)
	if err != nil {
		return nil, err
	}

	return &balance.Balance{
		Residual: residual,
		Balance:  cashes - residual,
	}, nil
}

func (s *Service) residual(ctx context.Context) (float64, error) {
	income, err := s.repository.Income(ctx)
	if err != nil {
		return 0, err
	}

	expense, err := s.repository.Expense(ctx)
	if err != nil {
		return 0, err
	}

	return income - expense, nil
}
