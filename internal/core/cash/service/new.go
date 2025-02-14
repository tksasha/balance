package service

import "github.com/tksasha/balance/internal/core/cash"

func New(repository cash.Repository) *Service {
	return &Service{
		repository: repository,
	}
}
