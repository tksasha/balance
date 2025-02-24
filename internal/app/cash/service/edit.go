package service

import (
	"context"

	"github.com/tksasha/balance/internal/app/cash"
)

func (s *Service) Edit(ctx context.Context, input string) (*cash.Cash, error) {
	return s.findByID(ctx, input)
}
