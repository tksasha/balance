package handlers_test

import (
	"database/sql"
	"testing"

	"github.com/tksasha/balance/internal/app/category"
	"github.com/tksasha/balance/internal/common/currency"
)

func cleanup(t *testing.T, db *sql.DB) {
	t.Helper()

	if _, err := db.ExecContext(t.Context(), `DELETE FROM categories`); err != nil {
		t.Fatal(err)
	}
}

func createCategory(t *testing.T, db *sql.DB, category *category.Category) {
	t.Helper()

	query := `INSERT INTO categories(name, slug, currency) VALUES(?, ?, ?)`

	result, err := db.ExecContext(
		t.Context(),
		query,
		category.Name,
		category.Slug,
		currency.Default,
	)
	if err != nil {
		t.Fatal(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		t.Fatal(err)
	}

	if id == 0 {
		t.Fatal("failed to create category")
	}
}
