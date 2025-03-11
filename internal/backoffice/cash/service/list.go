package service

import (
	"context"

	"github.com/tksasha/balance/internal/backoffice/cash"
)

func (s *Service) List(ctx context.Context) (cash.Cashes, error) {
	return s.repository.FindAll(ctx)
}
