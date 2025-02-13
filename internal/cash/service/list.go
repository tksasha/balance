package service

import (
	"context"

	"github.com/tksasha/balance/internal/cash"
)

func (s *Service) List(ctx context.Context) (cash.Cashes, error) {
	return s.repository.List(ctx)
}
