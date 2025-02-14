package service

import (
	"context"

	"github.com/tksasha/balance/internal/core/index"
)

type Service struct {
	repository index.Repository
}

func New(repository index.Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Balance(ctx context.Context) (float64, error) {
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
