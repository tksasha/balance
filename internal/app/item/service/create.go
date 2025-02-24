package service

import (
	"context"

	"github.com/tksasha/balance/internal/app/item"
	"github.com/tksasha/balance/pkg/validation"
)

func (s *Service) Create(ctx context.Context, request item.CreateRequest) (*item.Item, error) {
	item := &item.Item{
		Description: request.Description,
	}

	validate := validation.New()

	item.Date = validate.Date("date", request.Date)

	item.Formula, item.Sum = validate.Formula("sum", request.Formula)

	_ = validate.Presence("category", request.CategoryID)

	if err := s.setCategory(ctx, item, request.CategoryID, validate); err != nil {
		return nil, err
	}

	if validate.HasErrors() {
		return item, validate.Errors
	}

	if err := s.itemRepository.Create(ctx, item); err != nil {
		return nil, err
	}

	return item, nil
}
