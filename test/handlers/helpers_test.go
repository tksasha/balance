package handlers_test

import (
	"context"
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
		if _, err := db.Connection.ExecContext(ctx, "DELETE FROM categories"); err != nil {
			t.Fatalf("failed to truncate categories, error: %v", err)
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
