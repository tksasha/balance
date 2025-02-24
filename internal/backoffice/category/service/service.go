package service

import (
	"context"
	"errors"

	"github.com/tksasha/balance/internal/backoffice/category"
	"github.com/tksasha/balance/internal/common"
	"github.com/tksasha/validator"
)

type Service struct {
	*common.BaseService

	repository category.Repository
}

func New(baseService *common.BaseService, repository category.Repository) *Service {
	return &Service{
		BaseService: baseService,
		repository:  repository,
	}
}

func (s *Service) nameAlreadyExists(
	ctx context.Context,
	name string,
	categoryID int,
	validator *validator.Validator,
) error {
	if name == "" {
		return nil
	}

	category, err := s.repository.FindByName(ctx, name)

	if errors.Is(err, common.ErrRecordNotFound) {
		return nil
	}

	if err != nil {
		return err
	}

	if category.ID != categoryID {
		validator.Set("name", common.AlreadyExists)
	}

	return nil
}
