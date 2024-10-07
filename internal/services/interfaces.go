package services

import (
	"context"

	"github.com/tksasha/balance/internal/models"
)

type ItemGetter interface {
	GetItem(ctx context.Context, id string) (*models.Item, error)
}

type ItemUpdater interface {
	UpdateItem(ctx context.Context, item *models.Item) error
}

type ItemDeleter interface {
	DeleteItem(ctx context.Context, item *models.Item) error
}
