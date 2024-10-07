package services

import (
	"context"

	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/repositories"
)

func CreateItemService(ctx context.Context, itemCreator repositories.ItemCreator, item *models.Item) error {
	return itemCreator.CreateItem(ctx, item)
}
