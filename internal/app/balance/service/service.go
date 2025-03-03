package service

import (
	"github.com/tksasha/balance/internal/app/balance"
)

type Service struct {
	repository balance.Repository
}

func New(repository balance.Repository) *Service {
	return &Service{
		repository: repository,
	}
}
