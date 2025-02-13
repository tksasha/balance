package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/handlers"
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/pkg/currencies"
	"gotest.tools/v3/assert"
)

func TestCategoryDeleteHandler(t *testing.T) {
	ctx := t.Context()

	service, db := newCategoryService(ctx, t)
	defer func() {
		_ = db.Close()
	}()

	handler := handlers.NewCategoryDeleteHandler(service)

	mux := newMux(t, "DELETE /categories/{id}", handler)

	t.Run("responds 404 on category not found", func(t *testing.T) {
		cleanup(ctx, t)

		request := newDeleteRequest(ctx, t, "/categories/1348")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNotFound)
	})

	t.Run("responds 200 on category found", func(t *testing.T) {
		cleanup(ctx, t)

		categoryToCreate := &models.Category{
			ID:       1411,
			Name:     "Miscellaneous",
			Currency: currencies.EUR,
		}

		createCategory(ctx, t, categoryToCreate)

		request := newDeleteRequest(ctx, t, "/categories/1411?currency=eur")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)

		category := findCategoryByID(ctx, t, currencies.EUR, 1411)

		assert.Equal(t, category.ID, 1411)
		assert.Equal(t, category.Name, "Miscellaneous")
		assert.Equal(t, category.Currency, currencies.EUR)
		assert.Assert(t, !category.DeletedAt.Time.IsZero())
	})
}
