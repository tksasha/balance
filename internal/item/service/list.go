package service

import (
	"context"

	"github.com/tksasha/balance/internal/item"
)

func (s *Service) List(ctx context.Context) (item.Items, error) {
	return s.itemRepository.List(ctx)
}
