package service

import (
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
