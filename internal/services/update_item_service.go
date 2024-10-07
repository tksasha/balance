package services

import (
	"context"

	"github.com/tksasha/balance/internal/interfaces"
	"github.com/tksasha/balance/internal/models"
)

type UpdateItemService struct {
	itemUpdater interfaces.ItemUpdater
}

func NewUpdateItemService(itemUpdater interfaces.ItemUpdater) *UpdateItemService {
	return &UpdateItemService{
		itemUpdater: itemUpdater,
	}
}

func (s *UpdateItemService) UpdateItem(ctx context.Context, item *models.Item) error {
	return s.itemUpdater.UpdateItem(ctx, item)
}
