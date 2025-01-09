package handlers_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/handlers"
	"github.com/tksasha/balance/internal/middlewares"
	"github.com/tksasha/balance/internal/models"
	providers "github.com/tksasha/balance/internal/providers/test"
	"github.com/tksasha/balance/internal/repositories"
	"github.com/tksasha/balance/internal/services"
	"github.com/tksasha/balance/pkg/currencies"
	"gotest.tools/v3/assert"
)

func TestUpdateCategoryHandler(t *testing.T) { //nolint:funlen
	dbNameProvider := providers.NewDBNameProvider()

	db := db.Open(dbNameProvider)

	categoryRepository := repositories.NewCategoryRepository(db)

	categoryService := services.NewCategoryService(categoryRepository)

	middleware := middlewares.NewCurrencyMiddleware().Wrap(
		handlers.NewUpdateCategoryHandler(categoryService),
	)

	mux := http.NewServeMux()
	mux.Handle("PATCH /categories/{id}", middleware)

	ctx := context.Background()

	truncate := func() {
		if _, err := db.Connection.ExecContext(ctx, "DELETE FROM categories"); err != nil {
			t.Fatalf("failed to truncate categories, error: %v", err)
		}
	}

	t.Run("when category id is not a digit, it should respond with 404", func(t *testing.T) {
		t.Cleanup(truncate)

		request, err := http.NewRequestWithContext(ctx, http.MethodPatch, "/categories/abcd", nil)
		if err != nil {
			t.Fatalf("failed to build new request with context, error: %v", err)
		}

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNotFound)
	})

	t.Run("when category is not found by id, it should respond with 404", func(t *testing.T) {
		t.Cleanup(truncate)

		request, err := http.NewRequestWithContext(ctx, http.MethodPatch, "/categories/1141", nil)
		if err != nil {
			t.Fatalf("failed to build new request with context, error: %v", err)
		}

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNotFound)
	})

	t.Run("when category name is already exists, it should respond with 500", func(t *testing.T) {
		t.Cleanup(truncate)

		for id, name := range map[int]string{1151: "Heterogeneous", 11654: "Paraphernalia"} {
			if _, err := db.Connection.ExecContext(
				ctx,
				"INSERT INTO categories(id, name, currency) VALUES(?, ?, ?)",
				id,
				name,
				currencies.USD,
			); err != nil {
				t.Fatalf("failed to create category, error: %v", err)
			}
		}

		formData := url.Values{}
		formData.Add("name", "Paraphernalia")

		body := strings.NewReader(formData.Encode())

		request, err := http.NewRequestWithContext(ctx, http.MethodPatch, "/categories/1151?currency=usd", body)
		if err != nil {
			t.Fatalf("failed to build new request with context, error: %v", err)
		}

		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusInternalServerError)
	})

	t.Run("when category name is uniq, it should respond with 200", func(t *testing.T) {
		t.Cleanup(truncate)

		if _, err := db.Connection.ExecContext(
			ctx,
			"INSERT INTO categories(id, name, currency) VALUES(?, ?, ?)",
			1208,
			"Paraphernalia",
			currencies.USD,
		); err != nil {
			t.Fatalf("failed to create category, error: %v", err)
		}

		formData := url.Values{}
		formData.Add("name", "Heterogeneous")

		body := strings.NewReader(formData.Encode())

		request, err := http.NewRequestWithContext(ctx, http.MethodPatch, "/categories/1208?currency=usd", body)
		if err != nil {
			t.Fatalf("failed to build new request with context, error: %v", err)
		}

		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)

		category := &models.Category{}

		if err := db.Connection.QueryRowContext(
			ctx,
			"SELECT id, name, currency FROM categories WHERE id=? AND currency=?",
			1208,
			currencies.USD,
		).Scan(&category.ID, &category.Name, &category.Currency); err != nil {
			t.Fatalf("failed to get category, error: %v", err)
		}

		assert.Equal(t, category.ID, 1208)
		assert.Equal(t, category.Name, "Heterogeneous")
		assert.Equal(t, category.Currency, currencies.USD)
	})
}
