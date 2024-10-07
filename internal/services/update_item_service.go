package services

import (
	"context"

	"github.com/tksasha/balance/internal/repositories"
)

type UpdateItemService struct {
	itemGetter  ItemGetter
	itemUpdater repositories.ItemUpdater
}

func NewUpdateItemService(itemGetter ItemGetter, itemUpdater repositories.ItemUpdater) *UpdateItemService {
	return &UpdateItemService{
		itemGetter:  itemGetter,
		itemUpdater: itemUpdater,
	}
}

func (s *UpdateItemService) UpdateItem(ctx context.Context, id string) error {
	item, err := s.itemGetter.GetItem(ctx, id)
	if err != nil {
		return err
	}

	return s.itemUpdater.UpdateItem(ctx, item)
}
