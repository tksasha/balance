package services

import (
	"context"

	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/repositories"
)

type DeleteItemService struct {
	itemDeleter repositories.ItemDeleter
}

func NewDeleteItemService(itemDeleter repositories.ItemDeleter) *DeleteItemService {
	return &DeleteItemService{
		itemDeleter: itemDeleter,
	}
}

func (s *DeleteItemService) DeleteItem(ctx context.Context, item *models.Item) error {
	return s.itemDeleter.DeleteItem(ctx, item)
}
