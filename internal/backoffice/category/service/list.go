package service

import (
	"context"

	"github.com/tksasha/balance/internal/backoffice/category"
)

func (s *Service) List(ctx context.Context) (category.Categories, error) {
	return s.repository.FindAll(ctx)
}
