package service

import "context"

func (s *Service) Residual(ctx context.Context) (float64, error) {
	return s.residual(ctx)
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
