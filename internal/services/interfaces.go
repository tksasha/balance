package services

import (
	"context"

	"github.com/tksasha/balance/internal/models"
)

type ItemRepository interface {
	GetItems(ctx context.Context) (models.Items, error)
	Create(ctx context.Context, item *models.Item) error
	FindByID(ctx context.Context, id int) (*models.Item, error)
	Update(ctx context.Context, item *models.Item) error
	Delete(ctx context.Context, id int) error
}
