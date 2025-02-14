package service

import (
	"context"
	"errors"

	"github.com/tksasha/balance/internal/core/common/services"
	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/core/common"
	"github.com/tksasha/balance/pkg/validation"
)

type Service struct {
	repository category.Repository
}

func New(repository category.Repository) *Service {
	return &Service{
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
		validation.Set("name", services.AlreadyExists)
	}

	return nil
}
