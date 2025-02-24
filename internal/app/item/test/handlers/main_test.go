package handlers_test

import (
	"database/sql"
	"net/http"
	"testing"

	"github.com/tksasha/balance/internal/app/category"
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

func createCategory(t *testing.T, db *sql.DB, category *category.Category) {
	t.Helper()

	if _, err := db.ExecContext(
		t.Context(),
		"INSERT INTO categories(id, name, income, currency) VALUES(?, ?, ?, ?)",
		category.ID,
		category.Name,
		category.Income,
		category.Currency,
	); err != nil {
		t.Fatalf("failed to create category, error: %v", err)
	}
}
