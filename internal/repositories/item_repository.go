package repositories

import (
	"context"

	"github.com/tksasha/balance/internal/models"
)

type ItemRepository struct{}

func NewItemRepository() *ItemRepository {
	return &ItemRepository{}
}

func (r *ItemRepository) CreateItem(ctx context.Context, item *models.Item) error {
	return nil
}
