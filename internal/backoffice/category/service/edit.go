package service

import (
	"context"

	"github.com/tksasha/balance/internal/backoffice/category"
)

func (s *Service) Edit(ctx context.Context, input string) (*category.Category, error) {
	return s.findByID(ctx, input)
}
