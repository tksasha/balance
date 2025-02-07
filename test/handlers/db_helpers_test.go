package handlers_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/models"
	providers "github.com/tksasha/balance/internal/providers/test"
	"github.com/tksasha/balance/pkg/currencies"
)

func newDB(ctx context.Context, t *testing.T) *sql.DB {
	t.Helper()

	dbNameProvider := providers.NewDBNameProvider()

	return db.Open(ctx, dbNameProvider).Connection
}

func findCashByName(ctx context.Context, t *testing.T, currency currencies.Currency, name string) *models.Cash {
	t.Helper()

	ctx = currencyContext(ctx, t, currency)

	query := `
		SELECT id, name, formula, sum, currency, supercategory, favorite, deleted_at
		FROM cashes
		WHERE name=? AND currency=?
	`

	cash := &models.Cash{}

	db := newDB(ctx, t)
	defer func() {
		_ = db.Close()
	}()

	if err := db.
		QueryRowContext(ctx, query, name, currency).
		Scan(
			&cash.ID,
			&cash.Name,
			&cash.Formula,
			&cash.Sum,
			&cash.Currency,
			&cash.Supercategory,
			&cash.Favorite,
			&cash.DeletedAt,
		); err != nil {
		t.Fatalf("failed to find cash by name: %v", err)
	}

	return cash
}

func createCash(ctx context.Context, t *testing.T, cash *models.Cash) {
	t.Helper()

	db := newDB(ctx, t)
	defer func() {
		_ = db.Close()
	}()

	if _, err := db.ExecContext(
		ctx,
		"INSERT INTO cashes(id, currency, formula, sum, name, supercategory, favorite) VALUES(?, ?, ?, ?, ?, ?, ?)",
		cash.ID,
		cash.Currency,
		cash.Formula,
		cash.Sum,
		cash.Name,
		cash.Supercategory,
		cash.Favorite,
	); err != nil {
		t.Fatalf("failed to create cash: %v", err)
	}
}

func createCategory(ctx context.Context, t *testing.T, db *db.DB, category *models.Category) {
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

func findCategoryByName(ctx context.Context, t *testing.T, db *db.DB, name string) *models.Category {
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

func findCategoryByID(ctx context.Context, t *testing.T, db *db.DB, id int) *models.Category {
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

func cleanup(ctx context.Context, t *testing.T) {
	t.Helper()

	t.Cleanup(func() {
		db := newDB(ctx, t)
		defer func() {
			_ = db.Close()
		}()

		tables := []string{"items", "categories", "cashes"}

		for _, table := range tables {
			if _, err := db.ExecContext(ctx, "DELETE FROM "+table); err != nil { //nolint:gosec
				t.Fatalf("failed to truncate %s, error: %v", table, err)
			}
		}
	})
}

func createItem(ctx context.Context, t *testing.T, db *db.DB, item *models.Item) {
	t.Helper()

	query := `
		INSERT INTO items(id, date, formula, sum, category_id, category_name, description, currency)
		VALUES(?, ?, ?, ?, ?, ?, ?, ?)
	`

	if _, err := db.Connection.ExecContext(
		ctx,
		query,
		item.ID,
		item.Date,
		item.Formula,
		item.Sum,
		item.CategoryID,
		item.CategoryName,
		item.Description,
		item.Currency,
	); err != nil {
		t.Fatalf("failed to create item, error: %v", err)
	}
}

func findItemByDate(ctx context.Context, t *testing.T, db *db.DB, date string) *models.Item {
	t.Helper()

	currency, ok := ctx.Value(currencies.CurrencyContextValue{}).(currencies.Currency)
	if !ok {
		currency = currencies.DefaultCurrency
	}

	item := &models.Item{}

	query := `
		SELECT id, date, formula, sum, category_id, category_name, currency, description
		FROM items
		WHERE date=? AND currency=?
		ORDER BY created_at
		LIMIT 1
	`

	if err := db.Connection.
		QueryRowContext(ctx, query, date, currency).
		Scan(
			&item.ID,
			&item.Date,
			&item.Formula,
			&item.Sum,
			&item.CategoryID,
			&item.CategoryName,
			&item.Currency,
			&item.Description,
		); err != nil {
		t.Fatalf("failed to find item by date, error: %v", err)
	}

	return item
}
