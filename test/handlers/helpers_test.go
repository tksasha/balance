package handlers_test

import (
	"context"
	"database/sql"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/pkg/currencies"
)

type Params map[string]string

func createCategory(ctx context.Context, t *testing.T, db *db.DB, category *models.Category) {
	t.Helper()

	if _, err := db.Connection.ExecContext(
		ctx,
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

func findCategoryByName(ctx context.Context, t *testing.T, db *db.DB, name string) *models.Category {
	t.Helper()

	currency, ok := ctx.Value(currencies.CurrencyContextValue{}).(currencies.Currency)
	if !ok {
		currency = currencies.DefaultCurrency
	}

	query := `
		SELECT id, name, income, visible, currency, supercategory, deleted_at
		FROM categories
		WHERE name=? AND currency=?
	`

	category := &models.Category{}

	if err := db.Connection.
		QueryRowContext(ctx, query, name, currency).
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

func findCategoryByID(ctx context.Context, t *testing.T, db *db.DB, id int) *models.Category {
	t.Helper()

	currency, ok := ctx.Value(currencies.CurrencyContextValue{}).(currencies.Currency)
	if !ok {
		currency = currencies.DefaultCurrency
	}

	query := `
		SELECT id, name, income, visible, currency, supercategory, deleted_at
		FROM categories
		WHERE id=? AND currency=?
	`

	category := &models.Category{}

	if err := db.Connection.
		QueryRowContext(ctx, query, id, currency).
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

func findItemByDate(ctx context.Context, t *testing.T, db *sql.DB, date string) *models.Item {
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

	if err := db.
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

func usdContext(ctx context.Context, t *testing.T) context.Context {
	t.Helper()

	return context.WithValue(ctx, currencies.CurrencyContextValue{}, currencies.USD)
}

func eurContext(ctx context.Context, t *testing.T) context.Context {
	t.Helper()

	return context.WithValue(ctx, currencies.CurrencyContextValue{}, currencies.EUR)
}

func cleanup(ctx context.Context, t *testing.T, db *db.DB) {
	t.Helper()

	t.Cleanup(func() {
		tables := []string{"items", "categories"}

		for _, table := range tables {
			if _, err := db.Connection.ExecContext(ctx, "DELETE FROM "+table); err != nil { //nolint:gosec
				t.Fatalf("failed to truncate %s, error: %v", table, err)
			}
		}
	})
}

func newRequest(ctx context.Context, t *testing.T, method, endpoint string, params Params) *http.Request {
	t.Helper()

	formData := url.Values{}

	for name, value := range params {
		formData.Add(name, value)
	}

	body := strings.NewReader(formData.Encode())

	request, err := http.NewRequestWithContext(ctx, method, endpoint, body)
	if err != nil {
		t.Fatalf("failed to build new request with context, error: %v", err)
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return request
}

func newPostRequest(ctx context.Context, t *testing.T, endpoint string, params Params) *http.Request {
	t.Helper()

	return newRequest(ctx, t, http.MethodPost, endpoint, params)
}

func newPatchRequest(ctx context.Context, t *testing.T, endpoint string, params Params) *http.Request {
	t.Helper()

	return newRequest(ctx, t, http.MethodPatch, endpoint, params)
}

func newDeleteRequest(ctx context.Context, t *testing.T, endpoint string) *http.Request {
	t.Helper()

	return newRequest(ctx, t, http.MethodDelete, endpoint, nil)
}

func newGetRequest(ctx context.Context, t *testing.T, endpoint string) *http.Request {
	t.Helper()

	return newRequest(ctx, t, http.MethodGet, endpoint, nil)
}
