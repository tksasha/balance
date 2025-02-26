package handlers_test

import (
	"database/sql"
	"net/http"
	"testing"

	"github.com/tksasha/balance/internal/app/category"
	"github.com/tksasha/balance/internal/app/item"
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

func cleanup(t *testing.T, db *sql.DB) {
	t.Helper()

	if _, err := db.Exec(`DELETE FROM items; DELETE FROM categories`); err != nil {
		t.Fatal(err)
	}
}

func findItemByDate(t *testing.T, db *sql.DB, currency currency.Currency, date string) *item.Item {
	t.Helper()

	query := `
		SELECT
			id,
			currency,
			date,
			category_id,
			category_name,
			formula,
			sum,
			description
		FROM
			items
		where
			date = ?
			AND currency = ?
	`

	item := &item.Item{}

	if err := db.QueryRow(query, date, currency).Scan(
		&item.ID,
		&item.Currency,
		&item.Date,
		&item.CategoryID,
		&item.CategoryName,
		&item.Formula,
		&item.Sum,
		&item.Description,
	); err != nil {
		t.Fatal(err)
	}

	return item
}

func createItem(t *testing.T, db *sql.DB, item *item.Item) {
	t.Helper()

	query := `
		INSERT INTO
		    items (id, currency, date, category_id, sum, description)
		VALUES
		    (?, ?, ?, ?, ?, ?)
	`

	result, err := db.ExecContext(
		t.Context(),
		query,
		item.ID,
		item.Currency,
		item.Date,
		item.CategoryID,
		item.Sum,
		item.Description,
	)
	if err != nil {
		t.Fatal(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		t.Fatal(err)
	}

	if id == 0 {
		t.Fatal("failed to create item")
	}
}
