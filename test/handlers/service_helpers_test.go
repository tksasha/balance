package handlers_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/tksasha/balance/internal/handlers"
	"github.com/tksasha/balance/internal/repositories"
	"github.com/tksasha/balance/internal/services"
)

func newCashService(ctx context.Context, t *testing.T) (handlers.CashService, *sql.DB) {
	t.Helper()

	db := newDB(ctx, t)

	cashRepository := repositories.NewCashRepository(db)

	return services.NewCashService(cashRepository), db
}

func newCategoryService(ctx context.Context, t *testing.T) (handlers.CategoryService, *sql.DB) {
	t.Helper()

	db := newDB(ctx, t)

	categoryRepository := repositories.NewCategoryRepository(db)

	return services.NewCategoryService(categoryRepository), db
}
