package service

import (
	"context"

	"github.com/tksasha/balance/internal/app/item"
	"github.com/tksasha/balance/internal/common"
	"github.com/tksasha/validation"
)

func (s *Service) Update(ctx context.Context, request item.UpdateRequest) (*item.Item, error) {
	item, err := s.findByID(ctx, request.ID)
	if err != nil {
		return nil, err
	}

	validation := validation.New()

	item.Date = validation.Date("date", request.Date, common.DateFormat)

	item.Formula, item.Sum = validation.Formula("sum", request.Formula)

	_ = validation.Presence("category", request.CategoryID)

	if err := s.setCategory(ctx, item, request.CategoryID, validation); err != nil {
		return nil, err
	}

	item.Description = request.Description

	if validation.Errors.Exists() {
		return item, validation.Errors
	}

	if err := s.itemRepository.Update(ctx, item); err != nil {
		return nil, err
	}

	return item, nil
}
