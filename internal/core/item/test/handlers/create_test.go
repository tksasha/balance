package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/core/common/tests"
	"github.com/tksasha/balance/internal/core/item/handlers"
	"github.com/tksasha/balance/pkg/currencies"
	"gotest.tools/v3/assert"
)

func TestItemCreateHandler(t *testing.T) { //nolint:funlen
	ctx := t.Context()

	service, db := tests.NewItemService(ctx, t)
	defer func() {
		_ = db.Close()
	}()

	mux := tests.NewMux(t, "POST /items", handlers.NewCreateHandler(service))

	t.Run("responds 400 on parse form fails", func(t *testing.T) {
		request := tests.NewInvalidPostRequest(ctx, t, "/items")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("renders errors on invalid input", func(t *testing.T) {
		tests.Cleanup(ctx, t)

		params := tests.Params{"date": ""}

		request := tests.NewPostRequest(ctx, t, "/items", params)

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})

	t.Run("responds 200 on succcessful create", func(t *testing.T) {
		tests.Cleanup(ctx, t)

		categoryToCreate := &category.Category{
			ID:       1101,
			Name:     "Accoutrements",
			Currency: currencies.USD,
		}

		tests.CreateCategory(ctx, t, categoryToCreate)

		params := tests.Params{
			"date":        "2024-10-16",
			"formula":     "42.69+69.42",
			"category_id": "1101",
			"description": "paper clips, notebooks, and pens",
		}

		request := tests.NewPostRequest(ctx, t, "/items?currency=usd", params)

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)

		item := tests.FindItemByDate(ctx, t, currencies.USD, "2024-10-16")

		assert.Equal(t, item.Date.Format(time.DateOnly), "2024-10-16")
		assert.Equal(t, item.CategoryID, 1101)
		assert.Equal(t, item.CategoryName, "Accoutrements")
		assert.Equal(t, item.Currency, currencies.USD)
		assert.Equal(t, item.Formula, "42.69+69.42")
		assert.Equal(t, item.Sum, 112.11)
		assert.Equal(t, item.Description, "paper clips, notebooks, and pens")
	})
}
