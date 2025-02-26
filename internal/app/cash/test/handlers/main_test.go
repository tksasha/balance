package handlers_test

import (
	"database/sql"
	"net/http"
	"testing"

	"github.com/tksasha/balance/internal/app/cash"
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

func createCash(t *testing.T, db *sql.DB, cash *cash.Cash) {
	t.Helper()

	query := `
		INSERT INTO
		    cashes (
		        id,
		        currency,
		        formula,
		        sum,
		        name,
		        supercategory
		    )
		VALUES
		    (?, ?, ?, ?, ?, ?)
	`

	if _, err := db.ExecContext(
		t.Context(),
		query,
		cash.ID,
		cash.Currency,
		cash.Formula,
		cash.Sum,
		cash.Name,
		cash.Supercategory,
	); err != nil {
		t.Fatalf("failed to create cash: %v", err)
	}
}
