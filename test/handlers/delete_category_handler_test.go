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

	t.Run("when category id is not a digit, it should respond with 404", func(t *testing.T) {
		cleanup(ctx, t, db)

		request := newDeleteRequest(ctx, t, "/categories/abcd")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNotFound)
	})

	t.Run("when category is not found by id, it should respond with 404", func(t *testing.T) {
		cleanup(ctx, t, db)

		request := newDeleteRequest(ctx, t, "/categories/1348")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNotFound)
	})

	t.Run("when category is found, it should respond with 200", func(t *testing.T) {
		cleanup(ctx, t, db)

		createCategory(ctx, t, db,
			&models.Category{
				ID:       1411,
				Name:     "Miscellaneous",
				Currency: currencies.EUR,
			},
		)

		request := newDeleteRequest(ctx, t, "/categories/1411?currency=eur")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)

		category := findCategoryByID(eurContext(ctx, t), t, db, 1411)

		assert.Equal(t, category.ID, 1411)
		assert.Equal(t, category.Name, "Miscellaneous")
		assert.Equal(t, category.Currency, currencies.EUR)
		assert.Assert(t, !category.DeletedAt.Time.IsZero())
	})
}
