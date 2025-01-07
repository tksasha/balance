package handlers_test

import (
	"context"
	"io"
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
	is "gotest.tools/v3/assert/cmp"
)

func TestCreateCategoryHandlerTest(t *testing.T) { //nolint:funlen
	dbNameProvider := providers.NewDBNameProvider()

	db := db.Open(dbNameProvider)

	categoryRepository := repositories.NewCategoryRepository(db)

	categoryService := services.NewCategoryService(categoryRepository)

	middleware := middlewares.NewCurrencyMiddleware().Wrap(
		handlers.NewCreateCategoryHandler(categoryService),
	)

	mux := http.NewServeMux()
	mux.Handle("POST /categories", middleware)

	ctx := context.Background()

	truncate := func() {
		_, err := db.Connection.ExecContext(ctx, "DELETE FROM categories")
		if err != nil {
			t.Fatalf("failed to truncate categories, error: %v", err)
		}
	}

	t.Run("when input data is invalid, it should render validation errors", func(t *testing.T) {
		formData := url.Values{}
		formData.Set("name", "")

		body := strings.NewReader(formData.Encode())

		request, err := http.NewRequestWithContext(ctx, http.MethodPost, "/categories", body)
		if err != nil {
			t.Fatalf("failed to build request with context, error: %v", err)
		}

		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		response, err := io.ReadAll(recorder.Body)
		if err != nil {
			t.Fatalf("failed to read response body, error: %v", err)
		}

		assert.Equal(t, recorder.Code, http.StatusOK)
		assert.Assert(t, is.Contains(string(response), "name: is required"))
	})

	t.Run("when input data is valid, it should create category", func(t *testing.T) {
		t.Cleanup(truncate)

		formData := url.Values{}
		formData.Set("name", "Food")

		body := strings.NewReader(formData.Encode())

		request, err := http.NewRequestWithContext(ctx, http.MethodPost, "/categories?currency=eur", body)
		if err != nil {
			t.Fatalf("failed to build request with context, error: %v", err)
		}

		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		response, err := io.ReadAll(recorder.Body)
		if err != nil {
			t.Fatalf("failed to read response body, error: %v", err)
		}

		assert.Equal(t, recorder.Code, http.StatusOK)
		assert.Assert(t, is.Contains(string(response), "create category page"))

		category := &models.Category{}

		if err := db.Connection.QueryRowContext(
			ctx,
			"SELECT id, name, currency FROM categories WHERE name=? AND currency=?",
			"Food",
			currencies.EUR,
		).Scan(&category.ID, &category.Name, &category.Currency); err != nil {
			t.Fatalf("failed to find category by name, error: %v", err)
		}

		assert.Equal(t, category.ID, 1)
		assert.Equal(t, category.Name, "Food")
		assert.Equal(t, category.Currency, currencies.EUR)
	})
}
