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
	"github.com/tksasha/balance/test/testutils"
	"gotest.tools/v3/assert"
)

func TestEdit(t *testing.T) {
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

	t.Run("responds 404 on category not found", func(t *testing.T) {
		testutils.Cleanup(ctx, t, db)

		request := testutils.NewGetRequest(ctx, t, "/categories/1004/edit")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNotFound)
	})

	t.Run("responds 200 on category found", func(t *testing.T) {
		testutils.Cleanup(ctx, t, db)

		testutils.CreateCategory(ctx, t, db,
			&models.Category{
				ID:       1010,
				Name:     "Xenomorphic",
				Currency: currencies.EUR,
			},
		)

		request := testutils.NewGetRequest(ctx, t, "/categories/1010/edit?currency=eur")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)
	})
}
