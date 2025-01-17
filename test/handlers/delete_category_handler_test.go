package handlers_test

import (
	"context"
	"net/http"
	"net/http/httptest"
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

func TestDeleteCategoryHandler(t *testing.T) { //nolint:funlen
	dbNameProvider := providers.NewDBNameProvider()

	db := db.Open(dbNameProvider)

	categoryRepository := repositories.NewCategoryRepository(db)

	categoryService := services.NewCategoryService(categoryRepository)

	middleware := middlewares.NewCurrencyMiddleware().Wrap(
		handlers.NewDeleteCategoryHandler(categoryService),
	)

	mux := http.NewServeMux()
	mux.Handle("DELETE /categories/{id}", middleware)

	ctx := context.Background()

	truncate := func() {
		if _, err := db.Connection.ExecContext(ctx, "DELETE FROM categories"); err != nil {
			t.Fatalf("failed to truncate categories, error: %v", err)
		}
	}

	t.Run("when category id is not a digit, it should respond with 404", func(t *testing.T) {
		t.Cleanup(truncate)

		request, err := http.NewRequestWithContext(ctx, http.MethodDelete, "/categories/abcd", nil)
		if err != nil {
			t.Fatalf("failed to build new request with context, error: %v", err)
		}

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNotFound)
	})

	t.Run("when category is not found by id, it should respond with 404", func(t *testing.T) {
		t.Cleanup(truncate)

		request, err := http.NewRequestWithContext(ctx, http.MethodDelete, "/categories/1348", nil)
		if err != nil {
			t.Fatalf("failed to build new request with context, error: %v", err)
		}

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNotFound)
	})

	t.Run("when category is found, it should respond with 200", func(t *testing.T) {
		t.Cleanup(truncate)

		if _, err := db.Connection.ExecContext(
			ctx,
			"INSERT INTO categories(id, name, currency) VALUES(?, ?, ?)",
			1411,
			"Miscellaneous",
			currencies.EUR,
		); err != nil {
			t.Fatalf("failed to create category, error: %v", err)
		}

		request, err := http.NewRequestWithContext(ctx, http.MethodDelete, "/categories/1411?currency=eur", nil)
		if err != nil {
			t.Fatalf("failed to create new request with context, error: %v", err)
		}

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)

		category := &models.Category{}

		if err := db.Connection.QueryRowContext(
			ctx,
			"SELECT id, name, currency FROM categories WHERE id=?",
			1411,
		).Scan(&category.ID, &category.Name, &category.Currency); err != nil {
			t.Fatalf("failed to get category, error: %v", err)
		}

		assert.Equal(t, category.ID, 1411)
		assert.Equal(t, category.Name, "Miscellaneous")
		assert.Equal(t, category.Currency, currencies.EUR)
		// assert.Equal(t, category.DeletedAt, time.Now())
		{
		}
	})
}
