package services

import (
	"context"

	"github.com/tksasha/balance/internal/interfaces"
	"github.com/tksasha/balance/internal/models"
)

func CreateItemService(
	ctx context.Context,
	itemCreator interfaces.ItemCreator,
	item *models.Item,
) error {
	if err := itemCreator.CreateItem(ctx, item); err != nil {
		return err
	}

	return nil
}
