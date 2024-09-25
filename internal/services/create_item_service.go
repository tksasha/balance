package services

import (
	"context"

	"github.com/tksasha/balance/internal/repositories"
	"github.com/tksasha/balance/internal/requests"
)

func CreateItemService(
	ctx context.Context,
	repo *repositories.ItemRepository,
	req *requests.CreateItemRequest,
) error {
	if err := repo.CreateItem(ctx, req); err != nil {
		return err
	}

	return nil
}
