package interfaces

import (
	"context"

	"github.com/tksasha/balance/internal/models"
)

type ItemCreator interface {
	CreateItem(ctx context.Context, item *models.Item) error
}
