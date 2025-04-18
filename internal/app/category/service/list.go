package service

import (
	"context"

	"github.com/tksasha/balance/internal/app/category"
)

func (s *Service) List(ctx context.Context) (category.Categories, error) {
	return s.repository.FindAll(ctx)
}
