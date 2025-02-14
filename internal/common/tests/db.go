package tests

import (
	"context"
	"database/sql"
	"testing"

	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/core/cash"
	"github.com/tksasha/balance/internal/core/item"
	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/providers"
	"github.com/tksasha/balance/pkg/currencies"
)

func newDB(ctx context.Context, t *testing.T) *sql.DB {
	t.Helper()

	dbNameProvider := providers.NewDBNameProvider()

	return db.Open(ctx, dbNameProvider)
}

func Cleanup(ctx context.Context, t *testing.T) {
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

func FindCashByName(ctx context.Context, t *testing.T, currency currencies.Currency, name string) *cash.Cash {
	t.Helper()

	ctx = currencyContext(ctx, t, currency)

	query := `
		SELECT id, name, formula, sum, currency, supercategory, favorite, deleted_at
		FROM cashes
		WHERE name=? AND currency=?
	`

	cash := &cash.Cash{}

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

func CreateCash(ctx context.Context, t *testing.T, cash *cash.Cash) {
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

func FindCashByID(ctx context.Context, t *testing.T, currency currencies.Currency, id int) *cash.Cash {
	t.Helper()

	db := newDB(ctx, t)
	defer func() {
		_ = db.Close()
	}()

	ctx = currencyContext(ctx, t, currency)

	query := `
		SELECT id, currency, formula, sum, name, supercategory, favorite, deleted_at
		FROM cashes
		WHERE id=? AND currency=?
	`

	cash := &cash.Cash{}

	if err := db.
		QueryRowContext(ctx, query, id, currency).
		Scan(
			&cash.ID,
			&cash.Currency,
			&cash.Formula,
			&cash.Sum,
			&cash.Name,
			&cash.Supercategory,
			&cash.Favorite,
			&cash.DeletedAt,
		); err != nil {
		t.Fatalf("failed to find cash by id, error: %v", err)
	}

	return cash
}

func CreateCategory(ctx context.Context, t *testing.T, category *category.Category) {
	t.Helper()

	db := newDB(ctx, t)
	defer func() {
		_ = db.Close()
	}()

	if _, err := db.ExecContext(
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

func FindCategoryByID(ctx context.Context, t *testing.T, currency currencies.Currency, id int) *category.Category {
	t.Helper()

	db := newDB(ctx, t)
	defer func() {
		_ = db.Close()
	}()

	ctx = currencyContext(ctx, t, currency)

	query := `
		SELECT id, name, income, visible, currency, supercategory, deleted_at
		FROM categories
		WHERE id=? AND currency=?
	`

	category := &category.Category{}

	if err := db.
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

func FindCategoryByName(
	ctx context.Context,
	t *testing.T,
	currency currencies.Currency,
	name string,
) *category.Category {
	t.Helper()

	db := newDB(ctx, t)
	defer func() {
		_ = db.Close()
	}()

	ctx = currencyContext(ctx, t, currency)

	query := `
		SELECT id, name, income, visible, currency, supercategory, deleted_at
		FROM categories
		WHERE name=? AND currency=?
	`

	category := &category.Category{}

	if err := db.
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

func FindItemByDate(ctx context.Context, t *testing.T, currency currencies.Currency, date string) *item.Item {
	t.Helper()

	db := newDB(ctx, t)
	defer func() {
		_ = db.Close()
	}()

	ctx = currencyContext(ctx, t, currency)

	item := &item.Item{}

	query := `
		SELECT id, date, formula, sum, category_id, category_name, currency, description
		FROM items
		WHERE date=? AND currency=?
		ORDER BY created_at
		LIMIT 1
	`

	if err := db.
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

func CreateItem(ctx context.Context, t *testing.T, item *item.Item) {
	t.Helper()

	db := newDB(ctx, t)
	defer func() {
		_ = db.Close()
	}()

	query := `
		INSERT INTO items(id, date, formula, sum, category_id, category_name, description, currency)
		VALUES(?, ?, ?, ?, ?, ?, ?, ?)
	`

	if _, err := db.ExecContext(
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
