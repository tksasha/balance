package services

import (
	"context"

	"github.com/tksasha/balance/internal/decorators"
	"github.com/tksasha/balance/internal/models"
)

type ItemsGetter interface {
	GetItems(ctx context.Context) (*decorators.ItemsDecorator, error)
}

type ItemGetter interface {
	GetItem(ctx context.Context, id string) (*decorators.ItemDecorator, error)
}

type ItemUpdater interface {
	UpdateItem(ctx context.Context, item *models.Item) error
}

type ItemDeleter interface {
	DeleteItem(ctx context.Context, item *models.Item) error
}

type CategoriesGetter interface {
	GetCategories(ctx context.Context, currency int) (*decorators.CategoriesDecorator, error)
}
