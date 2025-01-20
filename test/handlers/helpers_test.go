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

func usdContext(ctx context.Context, t *testing.T) context.Context {
	t.Helper()

	return context.WithValue(ctx, currencies.CurrencyContextValue{}, currencies.USD)
}

func truncate(ctx context.Context, t *testing.T, db *db.DB) func() {
	t.Helper()

	return func() {
		if _, err := db.Connection.ExecContext(ctx, "DELETE FROM categories"); err != nil {
			t.Fatalf("failed to truncate categories, error: %v", err)
		}
	}
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

func newPatchRequest(ctx context.Context, t *testing.T, endpoint string, params Params) *http.Request {
	t.Helper()

	return newRequest(ctx, t, http.MethodPatch, endpoint, params)
}
