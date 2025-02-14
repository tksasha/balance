package service

import (
	"context"

	"github.com/tksasha/balance/internal/core/category"
)

func (s *Service) List(ctx context.Context) (category.Categories, error) {
	return s.repository.List(ctx)
}
