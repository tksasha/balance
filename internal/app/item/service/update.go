package service

import (
	"context"

	"github.com/tksasha/balance/internal/app/item"
	"github.com/tksasha/balance/pkg/validation"
)

func (s *Service) Update(ctx context.Context, request item.UpdateRequest) (*item.Item, error) {
	item, err := s.findByID(ctx, request.ID)
	if err != nil {
		return nil, err
	}

	validate := validation.New()

	item.Date = validate.Date("date", request.Date)

	item.Formula, item.Sum = validate.Formula("sum", request.Formula)

	_ = validate.Presence("category", request.CategoryID)

	if err := s.setCategory(ctx, item, request.CategoryID, validate); err != nil {
		return nil, err
	}

	item.Description = request.Description

	if validate.HasErrors() {
		return item, validate.Errors
	}

	if err := s.itemRepository.Update(ctx, item); err != nil {
		return nil, err
	}

	return item, nil
}
