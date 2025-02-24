package service

import (
	"context"
	"errors"
	"strconv"

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

func (s *Service) findByID(ctx context.Context, input string) (*category.Category, error) {
	id, err := strconv.Atoi(input)
	if err != nil {
		return nil, common.ErrResourceNotFound
	}

	category, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return nil, s.MapError(err)
	}

	return category, nil
}
