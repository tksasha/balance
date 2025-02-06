package testutils

import (
	"context"
	"testing"

	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/pkg/currencies"
)

func CreateCategory(ctx context.Context, t *testing.T, db *db.DB, category *models.Category) {
	t.Helper()

	if _, err := db.Connection.ExecContext(
		ctx,
		"INSERT INTO categories(id, name, income, visible, currency, supercategory) VALUES(?, ?, ?, ?, ?, ?)",
		category.ID,
		category.Name,
		category.Income,
		category.Visible,
		category.Currency,
		category.Supercategory,
	); err != nil {
		t.Fatalf("failed to create category, error: %v", err)
	}
}

func FindCategoryByName(ctx context.Context, t *testing.T, db *db.DB, name string) *models.Category {
	t.Helper()

	currency, ok := ctx.Value(currencies.CurrencyContextValue{}).(currencies.Currency)
	if !ok {
		currency = currencies.DefaultCurrency
	}

	query := `
		SELECT id, name, income, visible, currency, supercategory, deleted_at
		FROM categories
		WHERE name=? AND currency=?
	`

	category := &models.Category{}

	if err := db.Connection.
		QueryRowContext(ctx, query, name, currency).
		Scan(
			&category.ID,
			&category.Name,
			&category.Income,
			&category.Visible,
			&category.Currency,
			&category.Supercategory,
			&category.DeletedAt,
		); err != nil {
		t.Fatalf("failed to find category by name, error: %v", err)
	}

	return category
}

func FindCategoryByID(ctx context.Context, t *testing.T, db *db.DB, id int) *models.Category {
	t.Helper()

	currency, ok := ctx.Value(currencies.CurrencyContextValue{}).(currencies.Currency)
	if !ok {
		currency = currencies.DefaultCurrency
	}

	query := `
		SELECT id, name, income, visible, currency, supercategory, deleted_at
		FROM categories
		WHERE id=? AND currency=?
	`

	category := &models.Category{}

	if err := db.Connection.
		QueryRowContext(ctx, query, id, currency).
		Scan(
			&category.ID,
			&category.Name,
			&category.Income,
			&category.Visible,
			&category.Currency,
			&category.Supercategory,
			&category.DeletedAt,
		); err != nil {
		t.Fatalf("failed to find category by id, error: %v", err)
	}

	return category
}
