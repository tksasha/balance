package services

import (
	"context"

	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/repositories"
)

type UpdateItemService struct {
	itemUpdater repositories.ItemUpdater
}

func NewUpdateItemService(itemUpdater repositories.ItemUpdater) *UpdateItemService {
	return &UpdateItemService{
		itemUpdater: itemUpdater,
	}
}

func (s *UpdateItemService) UpdateItem(ctx context.Context, item *models.Item) error {
	return s.itemUpdater.UpdateItem(ctx, item)
}
