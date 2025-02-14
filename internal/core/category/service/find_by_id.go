package service

import (
	"context"
	"strconv"

	"github.com/tksasha/balance/internal/apperrors"
	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/common/services"
)

func (s *Service) FindByID(ctx context.Context, input string) (*category.Category, error) {
	return s.findByID(ctx, input)
}

func (s *Service) findByID(ctx context.Context, input string) (*category.Category, error) {
	id, err := strconv.Atoi(input)
	if err != nil {
		return nil, apperrors.ErrResourceNotFound
	}

	category, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return nil, services.E(err)
	}

	return category, nil
}
