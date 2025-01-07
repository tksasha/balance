package handlers_test

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/handlers"
	"github.com/tksasha/balance/internal/middlewares"
	"github.com/tksasha/balance/internal/models"
	providers "github.com/tksasha/balance/internal/providers/test"
	"github.com/tksasha/balance/internal/repositories"
	"github.com/tksasha/balance/internal/services"
	"gotest.tools/v3/assert"
)

func TestGetCategoriesHandler_GetCategories(t *testing.T) { //nolint:funlen
	db := db.Open(
		providers.NewDBNameProvider(),
	)

	categoryRepository := repositories.NewCategoryRepository(db)

	categoryService := services.NewCategoryService(categoryRepository)

	middleware := middlewares.NewCurrencyMiddleware().Wrap(
		handlers.NewGetCategoriesHandler(categoryService),
	)

	mux := http.NewServeMux()
	mux.Handle("GET /categories", middleware)

	ctx := context.Background()

	truncate := func() {
		_, err := db.Connection.ExecContext(ctx, "DELETE FROM categories")
		if err != nil {
			t.Fatalf("failed to truncate categories, error: %v", err)
		}
	}

	t.Run("when there no categories, it should render empty widget", func(t *testing.T) {
		t.Cleanup(truncate)

		request, err := http.NewRequestWithContext(ctx, http.MethodGet, "/categories?currency=eur", nil)
		if err != nil {
			t.Fatalf("failed to build new request with context, error: %v", err)
		}

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)
	})

	t.Run("when there are categories, it should render widget with those categories", func(t *testing.T) {
		t.Cleanup(truncate)

		for _, name := range []string{"category one", "category two"} {
			if _, err := db.Connection.ExecContext(
				ctx,
				"INSERT INTO categories(name, currency) VALUES(?, ?)",
				name,
				models.EUR,
			); err != nil {
				t.Fatalf("failed to create category with name %s, error: %v", name, err)
			}
		}

		request, err := http.NewRequestWithContext(ctx, http.MethodGet, "/categories?currency=eur", nil)
		if err != nil {
			t.Fatalf("failed to build new request with context, error: %v", err)
		}

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		response, err := io.ReadAll(recorder.Body)
		if err != nil {
			t.Fatalf("failed to read response body, error: %v", err)
		}

		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Assert(t, strings.Contains(string(response), "category one"))
		assert.Assert(t, strings.Contains(string(response), "category two"))
	})
}
