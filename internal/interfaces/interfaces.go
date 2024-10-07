package interfaces

import (
	"context"

	"github.com/tksasha/balance/internal/models"
)

type ItemCreator interface {
	CreateItem(ctx context.Context, item *models.Item) error
}

type ItemGetter interface {
	GetItem(ctx context.Context, id int) (*models.Item, error)
}

type ItemUpdater interface {
	UpdateItem(ctx context.Context, item *models.Item) error
}
