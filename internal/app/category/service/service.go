package service

import (
	"github.com/tksasha/balance/internal/app/category"
	"github.com/tksasha/balance/internal/common"
)

type Service struct {
	*common.BaseService

	repository category.Repository
}

func New(
	baseService *common.BaseService,
	repository category.Repository,
) *Service {
	return &Service{
		BaseService: baseService,
		repository:  repository,
	}
}
