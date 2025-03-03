package service

import (
	"context"
)

func (s *Service) Balance(ctx context.Context) (float64, float64, error) {
	residual, err := s.residual(ctx)
	if err != nil {
		return 0.0, 0.0, err
	}

	cashes, err := s.repository.Cashes(ctx)
	if err != nil {
		return 0.0, 0.0, err
	}

	return residual, cashes - residual, nil
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
