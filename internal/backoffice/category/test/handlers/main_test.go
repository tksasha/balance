package handlers_test

import (
	"database/sql"
	"net/http"
	"testing"

	"github.com/tksasha/balance/internal/backoffice/category"
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
		if _, err := db.Exec(`DELETE FROM categories`); err != nil {
			t.Fatal(err)
		}
	})
}

func createCategory(t *testing.T, db *sql.DB, category *category.Category) {
	t.Helper()

	if _, err := db.ExecContext(
		t.Context(),
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

func findCategoryByID(t *testing.T, db *sql.DB, currency currency.Currency, id int) *category.Category {
	t.Helper()

	query := `
		SELECT
			id,
			name,
			income,
			visible,
			currency,
			supercategory,
			deleted_at
		FROM
			categories
		WHERE
			id=?
			AND currency=?
	`

	category := &category.Category{}

	if err := db.
		QueryRowContext(t.Context(), query, id, currency).
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

func findCategoryByName(t *testing.T, db *sql.DB, currency currency.Currency, name string) *category.Category {
	t.Helper()

	query := `
		SELECT id, name, income, visible, currency, supercategory, deleted_at
		FROM categories
		WHERE name=? AND currency=?
	`

	category := &category.Category{}

	if err := db.
		QueryRowContext(t.Context(), query, name, currency).
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
