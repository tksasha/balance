package service

import (
	"context"

	"github.com/tksasha/balance/internal/app/item"
)

func (s *Service) Edit(ctx context.Context, input string) (*item.Item, error) {
	return s.findByID(ctx, input)
}
