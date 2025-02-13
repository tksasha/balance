package handlers_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/tksasha/balance/internal/category/repository"
	"github.com/tksasha/balance/internal/handlers"
	"github.com/tksasha/balance/internal/repositories"
	"github.com/tksasha/balance/internal/services"
)

func newItemService(ctx context.Context, t *testing.T) (handlers.ItemService, *sql.DB) {
	t.Helper()

	db := newDB(ctx, t)

	itemRepository := repositories.NewItemRepository(db)

	categoryRepository := repository.New(db)

	return services.NewItemService(itemRepository, categoryRepository), db
}
