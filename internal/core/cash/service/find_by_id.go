package service

import (
	"context"
	"errors"
	"strconv"

	"github.com/tksasha/balance/internal/core/cash"
	"github.com/tksasha/balance/internal/core/common"
)

func (s *Service) FindByID(ctx context.Context, input string) (*cash.Cash, error) {
	return s.findByID(ctx, input)
}

func (s *Service) findByID(ctx context.Context, input string) (*cash.Cash, error) {
	id, err := strconv.Atoi(input)
	if err != nil {
		return nil, common.ErrResourceNotFound
	}

	cash, err := s.repository.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, common.ErrRecordNotFound) {
			return nil, common.ErrResourceNotFound
		}

		return nil, err
	}

	return cash, nil
}
