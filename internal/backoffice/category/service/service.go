package service

import (
	"context"
	"errors"
	"strconv"

	"github.com/tksasha/balance/internal/backoffice/category"
	"github.com/tksasha/balance/internal/common"
	"github.com/tksasha/balance/internal/common/service"
	"github.com/tksasha/validation"
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

func (s *Service) nameAlreadyExists(
	ctx context.Context,
	name string,
	categoryID int,
	validation *validation.Validation,
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
		validation.Errors.Set("name", common.AlreadyExists)
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
