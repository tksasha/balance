package handlers

import (
	"context"

	"github.com/tksasha/balance/internal/models"
)

type ItemService interface {
	GetItems(ctx context.Context) (models.Items, error)
	GetItem(ctx context.Context, input string) (*models.Item, error)
	UpdateItem(ctx context.Context, item *models.Item) error
	DeleteItem(ctx context.Context, input string) error
	CreateItem(ctx context.Context, item *models.Item) error
}

type CategoryService interface {
	GetCategories(ctx context.Context) (models.Categories, error)
	Create(ctx context.Context, category *models.Category) error
}
