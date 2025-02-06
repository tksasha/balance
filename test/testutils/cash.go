package testutils

import (
	"context"
	"testing"

	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/models"
)

func FindCashByName(ctx context.Context, t *testing.T, db *db.DB, name string) *models.Cash {
	t.Helper()

	currency := getCurrency(ctx)

	query := `
		SELECT id, name, formula, sum, currency, supercategory, favorite, deleted_at
		FROM cashes
		WHERE name=? AND currency=?
	`

	cash := &models.Cash{}

	if err := db.Connection.
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
		t.Fatalf("failed to find cash by name, error: %v", err)
	}

	return cash
}
