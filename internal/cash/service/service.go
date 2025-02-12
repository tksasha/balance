package service

import (
	"context"

	"github.com/tksasha/balance/internal/cash"
)

type Service struct {
	repository cash.Repository
}

func New(repository cash.Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) List(ctx context.Context) (cash.Cashes, error) {
	return s.repository.List(ctx)
}
