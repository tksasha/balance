package service

import (
	"github.com/tksasha/balance/internal/app/category"
	"github.com/tksasha/balance/internal/common/service"
)

type Service struct {
	*service.Service

	repository category.Repository
}

func New(repository category.Repository) *Service {
	return &Service{
		Service:    service.New(),
		repository: repository,
	}
}
