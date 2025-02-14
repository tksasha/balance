package service

import (
	"context"
	"strconv"

	"github.com/tksasha/balance/internal/apperrors"
	"github.com/tksasha/balance/internal/common/services"
	"github.com/tksasha/balance/internal/item"
)

func (s *Service) FindByID(ctx context.Context, input string) (*item.Item, error) {
	return s.findByID(ctx, input)
}

func (s *Service) findByID(ctx context.Context, input string) (*item.Item, error) {
	id, err := strconv.Atoi(input)
	if err != nil || id <= 0 {
		return nil, apperrors.ErrResourceNotFound
	}

	item, err := s.itemRepository.FindByID(ctx, id)
	if err != nil {
		return nil, services.E(err)
	}

	return item, nil
}
