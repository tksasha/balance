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
	"github.com/tksasha/balance/pkg/currencies"
	"gotest.tools/v3/assert"
)

func TestCategoryListHandler(t *testing.T) {
	db := db.Open(
		providers.NewDBNameProvider(),
	)

	categoryRepository := repositories.NewCategoryRepository(db)

	categoryService := services.NewCategoryService(categoryRepository)

	middleware := middlewares.NewCurrencyMiddleware().Wrap(
		handlers.NewCategoryListHandler(categoryService),
	)

	mux := http.NewServeMux()
	mux.Handle("GET /categories", middleware)

	ctx := context.Background()

	t.Run("responds 200 on no categories found", func(t *testing.T) {
		cleanup(ctx, t, db)

		request := newGetRequest(ctx, t, "/categories?currency=eur")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)
	})

	t.Run("responds 200 on categories found", func(t *testing.T) {
		cleanup(ctx, t, db)

		for id, name := range map[int]string{1: "category one", 2: "category two"} {
			createCategory(ctx, t, db,
				&models.Category{
					ID:       id,
					Name:     name,
					Currency: currencies.EUR,
					Visible:  true,
				},
			)
		}

		request := newGetRequest(ctx, t, "/categories?currency=eur")

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
