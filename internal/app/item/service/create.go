package service

import (
	"context"

	"github.com/tksasha/balance/internal/app/item"
	"github.com/tksasha/balance/internal/common"
	"github.com/tksasha/validation"
)

func (s *Service) Create(ctx context.Context, request item.CreateRequest) (*item.Item, error) {
	item := &item.Item{
		Description: request.Description,
	}

	validation := validation.New()

	item.Date = validation.Date("date", request.Date, common.DateFormat)

	item.Formula, item.Sum = validation.Formula("sum", request.Formula)

	_ = validation.Presence("category", request.CategoryID)

	if err := s.setCategory(ctx, item, request.CategoryID, validation); err != nil {
		return nil, err
	}

	if validation.Errors.Exists() {
		return item, validation.Errors
	}

	if err := s.itemRepository.Create(ctx, item); err != nil {
		return nil, err
	}

	return item, nil
}
