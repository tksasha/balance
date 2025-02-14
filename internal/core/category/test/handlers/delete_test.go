package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/core/category/handlers"
	"github.com/tksasha/balance/internal/core/common/tests"
	"github.com/tksasha/balance/pkg/currencies"
	"gotest.tools/v3/assert"
)

func TestCategoryDeleteHandler(t *testing.T) {
	ctx := t.Context()

	service, db := tests.NewCategoryService(ctx, t)
	defer func() {
		_ = db.Close()
	}()

	handler := handlers.NewDeleteHandler(service)

	mux := tests.NewMux(t, "DELETE /categories/{id}", handler)

	t.Run("responds 404 when category not found", func(t *testing.T) {
		tests.Cleanup(ctx, t)

		request := tests.NewDeleteRequest(ctx, t, "/categories/1348")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNotFound)
	})

	t.Run("responds 204 when category deleted", func(t *testing.T) {
		tests.Cleanup(ctx, t)

		categoryToCreate := &category.Category{
			ID:       1411,
			Name:     "Miscellaneous",
			Currency: currencies.EUR,
		}

		tests.CreateCategory(ctx, t, categoryToCreate)

		request := tests.NewDeleteRequest(ctx, t, "/categories/1411?currency=eur")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNoContent)

		category := tests.FindCategoryByID(ctx, t, currencies.EUR, 1411)

		assert.Equal(t, category.ID, 1411)
		assert.Equal(t, category.Name, "Miscellaneous")
		assert.Equal(t, category.Currency, currencies.EUR)
		assert.Assert(t, !category.DeletedAt.Time.IsZero())
	})
}
