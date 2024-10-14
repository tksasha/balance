package services

import (
	"context"

	"github.com/tksasha/balance/internal/decorators"
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/repositories"
)

type GetItemsService struct {
	itemsGetter repositories.ItemsGetter
}

func NewGetItemsService(itemsGetter repositories.ItemsGetter) ItemsGetter {
	return &GetItemsService{
		itemsGetter: itemsGetter,
	}
}

func (s *GetItemsService) GetItems(ctx context.Context, currency *models.Currency) (*decorators.ItemsDecorator, error) {
	items, err := s.itemsGetter.GetItems(ctx, currency)
	if err != nil {
		return nil, err
	}

	return decorators.NewItemsDecorator(items), nil
}
