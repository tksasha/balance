package repository_test

import (
	"database/sql"
	"testing"
	"time"

	"github.com/tksasha/balance/internal/app/category"
	"github.com/tksasha/balance/internal/app/item"
	"github.com/tksasha/balance/internal/common/currency"
)

func createCategory(t *testing.T, db *sql.DB, category *category.Category) {
	t.Helper()

	query := `INSERT INTO categories(id, name, slug, income, supercategory) VALUES(?, ?, ?, ?, ?)`

	result, err := db.ExecContext(
		t.Context(),
		query,
		category.ID,
		category.Name,
		category.Slug,
		category.Income,
		category.Supercategory,
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

func findCategoryByID(t *testing.T, db *sql.DB, id int) *category.Category {
	t.Helper()

	query := `SELECT name, slug FROM categories WHERE id = ?`

	category := &category.Category{}

	if err := db.QueryRowContext(t.Context(), query, id).Scan(
		&category.Name,
		&category.Slug,
	); err != nil {
		t.Fatal(err)
	}

	return category
}

func createItem(t *testing.T, db *sql.DB, item *item.Item) {
	t.Helper()

	category := findCategoryByID(t, db, item.CategoryID)

	query := `
		INSERT INTO
			items(currency, date, category_id, category_name, category_slug, sum)
		VALUES(?, ?, ?, ?, ?, ?)
	`

	result, err := db.ExecContext(
		t.Context(),
		query,
		currency.Default,
		item.Date,
		item.CategoryID,
		category.Name,
		category.Slug,
		item.Sum,
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

func date(t *testing.T, value string) time.Time {
	t.Helper()

	date, err := time.Parse(time.DateOnly, value)
	if err != nil {
		t.Fatal(err)
	}

	return date
}

func cleanup(t *testing.T, db *sql.DB) {
	t.Helper()

	if _, err := db.ExecContext(t.Context(), `DELETE FROM items; DELETE FROM categories;`); err != nil {
		t.Fatal(err)
	}
}
