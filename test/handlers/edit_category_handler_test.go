package handlers_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/handlers"
	"github.com/tksasha/balance/internal/middlewares"
	providers "github.com/tksasha/balance/internal/providers/test"
	"github.com/tksasha/balance/internal/repositories"
	"github.com/tksasha/balance/internal/services"
	"github.com/tksasha/balance/pkg/currencies"
	"gotest.tools/v3/assert"
)

func TestEditCategoryHandler(t *testing.T) { //nolint:funlen
	dbNameProvider := providers.NewDBNameProvider()

	db := db.Open(dbNameProvider)

	categoryRepository := repositories.NewCategoryRepository(db)

	categoryService := services.NewCategoryService(categoryRepository)

	middleware := middlewares.NewCurrencyMiddleware().Wrap(
		handlers.NewEditCategoryHandler(categoryService),
	)

	mux := http.NewServeMux()
	mux.Handle("GET /categories/{id}/edit", middleware)

	ctx := context.Background()

	truncate := func() {
		if _, err := db.Connection.ExecContext(ctx, "DELETE FROM categories"); err != nil {
			t.Fatalf("failed to truncate categories")
		}
	}

	t.Run("when category id is not a digit, it should respond with 404", func(t *testing.T) {
		t.Cleanup(truncate)

		request, err := http.NewRequestWithContext(ctx, http.MethodGet, "/categories/abcd/edit", nil)
		if err != nil {
			t.Fatalf("failed to build new request with context, error: %v", err)
		}

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNotFound)
	})

	t.Run("when category was not found by id, it should respond with 404", func(t *testing.T) {
		t.Cleanup(truncate)

		request, err := http.NewRequestWithContext(ctx, http.MethodGet, "/categories/1004/edit", nil)
		if err != nil {
			t.Fatalf("failed to build new request with context, error: %v", err)
		}

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNotFound)
	})

	t.Run("when category was found by id, it should respond with 200", func(t *testing.T) {
		t.Cleanup(truncate)

		if _, err := db.Connection.ExecContext(
			ctx,
			"INSERT INTO categories (id, name, currency) VALUES (?, ?, ?)",
			1010,
			"Xenomorphic",
			currencies.EUR,
		); err != nil {
			t.Fatalf("failed to create category, error: %v", err)
		}

		request, err := http.NewRequestWithContext(ctx, http.MethodGet, "/categories/1010/edit?currency=eur", nil)
		if err != nil {
			t.Fatalf("failed to build new request with context, error: %v", err)
		}

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)
	})
}
