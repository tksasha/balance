package service

import (
	"github.com/tksasha/balance/internal/backoffice/category"
)

type Service struct {
	repository category.Repository
}

func New(repository category.Repository) *Service {
	return &Service{
		repository: repository,
	}
}
