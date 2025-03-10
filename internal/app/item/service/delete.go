package service

import (
	"context"

	"github.com/tksasha/balance/internal/app/item"
)

func (s *Service) Delete(ctx context.Context, input string) (*item.Item, error) {
	item, err := s.findByID(ctx, input)
	if err != nil {
		return nil, err
	}

	if err := s.itemRepository.Delete(ctx, item.ID); err != nil {
		return nil, s.MapError(err)
	}

	return item, nil
}
