package testutils

import (
	"context"
	"testing"

	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/pkg/currencies"
)

func CreateItem(ctx context.Context, t *testing.T, db *db.DB, item *models.Item) {
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

func FindItemByDate(ctx context.Context, t *testing.T, db *db.DB, date string) *models.Item {
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
