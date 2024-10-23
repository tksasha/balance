package repositories

import (
	"context"

	"github.com/tksasha/balance/internal/models"
)

type ItemRepository interface {
	GetItems(ctx context.Context) (models.Items, error)
	CreateItem(ctx context.Context, item *models.Item) error
	GetItem(ctx context.Context, id int) (*models.Item, error)
	UpdateItem(ctx context.Context, item *models.Item) error
	DeleteItem(ctx context.Context, id int) error
}

type itemRepository struct{}

func NewItemRepository() ItemRepository {
	return &itemRepository{}
}

func (r *itemRepository) GetItems(ctx context.Context) (models.Items, error) {
	return models.Items{}, nil
}

func (r *itemRepository) CreateItem(ctx context.Context, item *models.Item) error {
	return nil
}

func (r *itemRepository) GetItem(ctx context.Context, id int) (*models.Item, error) {
	return &models.Item{}, nil
}

func (r *itemRepository) UpdateItem(ctx context.Context, item *models.Item) error {
	return nil
}

func (r *itemRepository) DeleteItem(ctx context.Context, id int) error {
	return nil
}
