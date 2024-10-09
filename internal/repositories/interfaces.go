package repositories

import (
	"context"

	"github.com/tksasha/balance/internal/models"
)

type ItemsGetter interface {
	GetItems(ctx context.Context, currency models.Currency) ([]*models.Item, error)
}

type ItemCreator interface {
	CreateItem(ctx context.Context, item *models.Item) error
}

type ItemGetter interface {
	GetItem(ctx context.Context, id int) (*models.Item, error)
}

type ItemUpdater interface {
	UpdateItem(ctx context.Context, item *models.Item) error
}

type ItemDeleter interface {
	DeleteItem(ctx context.Context, item *models.Item) error
}

type CategoriesGetter interface {
	GetCategories(ctx context.Context, currency int) ([]*models.Category, error)
}
