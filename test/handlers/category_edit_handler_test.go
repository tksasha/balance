package handlers_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/handlers"
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/pkg/currencies"
	"gotest.tools/v3/assert"
)

func TestCategoryEditHandler(t *testing.T) {
	ctx := context.Background()

	service, db := newCategoryService(ctx, t)
	defer func() {
		_ = db.Close()
	}()

	mux := newMux(t, "GET /categories/{id}/edit", handlers.NewCategoryEditHandler(service))

	t.Run("responds 404 on category not found", func(t *testing.T) {
		cleanup(ctx, t)

		request := newGetRequest(ctx, t, "/categories/1004/edit")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNotFound)
	})

	t.Run("responds 200 on category found", func(t *testing.T) {
		cleanup(ctx, t)

		categoryToCreate := &models.Category{
			ID:       1010,
			Name:     "Xenomorphic",
			Currency: currencies.EUR,
		}

		createCategory(ctx, t, categoryToCreate)

		request := newGetRequest(ctx, t, "/categories/1010/edit?currency=eur")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)
	})
}
