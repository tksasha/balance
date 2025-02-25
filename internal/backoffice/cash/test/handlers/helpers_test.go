package handlers_test

import (
	"database/sql"
	"net/http"
	"testing"

	"github.com/tksasha/balance/internal/app/cash"
	"github.com/tksasha/balance/internal/common/currency"
	"github.com/tksasha/balance/internal/server/middlewares"
)

func mux(t *testing.T, pattern string, handler http.Handler) *http.ServeMux {
	t.Helper()

	for _, middleware := range middlewares.New() {
		handler = middleware.Wrap(handler)
	}

	mux := http.NewServeMux()

	mux.Handle(pattern, handler)

	return mux
}

func cleanup(t *testing.T, db *sql.DB) {
	t.Helper()

	t.Cleanup(func() {
		if _, err := db.Exec(`DELETE FROM cashes`); err != nil {
			t.Fatal(err)
		}
	})
}

func findCashByName(t *testing.T, db *sql.DB, currency currency.Currency, name string) *cash.Cash {
	t.Helper()

	query := `
		SELECT
			id,
			name,
			formula,
			sum,
			currency,
			supercategory,
			favorite,
			deleted_at
		FROM
			cashes
		WHERE
			name=?
			AND currency=?
	`

	cash := &cash.Cash{}

	if err := db.
		QueryRowContext(t.Context(), query, name, currency).
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
