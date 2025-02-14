package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/core/category/handlers"
	"github.com/tksasha/balance/internal/common/tests"
	"github.com/tksasha/balance/pkg/currencies"
	"gotest.tools/v3/assert"
)

func TestCategoryEditHandler(t *testing.T) {
	ctx := t.Context()

	service, db := tests.NewCategoryService(ctx, t)
	defer func() {
		_ = db.Close()
	}()

	mux := tests.NewMux(t, "GET /categories/{id}/edit", handlers.NewEditHandler(service))

	t.Run("responds 404 on category not found", func(t *testing.T) {
		tests.Cleanup(ctx, t)

		request := tests.NewGetRequest(ctx, t, "/categories/1004/edit")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNotFound)
	})

	t.Run("responds 200 on category found", func(t *testing.T) {
		tests.Cleanup(ctx, t)

		categoryToCreate := &category.Category{
			ID:       1010,
			Name:     "Xenomorphic",
			Currency: currencies.EUR,
		}

		tests.CreateCategory(ctx, t, categoryToCreate)

		request := tests.NewGetRequest(ctx, t, "/categories/1010/edit?currency=eur")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)
	})
}
