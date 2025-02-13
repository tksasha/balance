package testutils

import (
	"context"
	"database/sql"
	"testing"

	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/models"
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

func FindCashByName(ctx context.Context, t *testing.T, currency currencies.Currency, name string) *models.Cash {
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
