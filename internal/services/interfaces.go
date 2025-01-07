package services

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

type CategoryRepository interface {
	GetCategories(ctx context.Context) (models.Categories, error)
	Create(ctx context.Context, category *models.Category) error
	FindByName(ctx context.Context, name string) (*models.Category, error)
}
