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

func TestItemCreateHandler(t *testing.T) { //nolint:funlen
	ctx := context.Background()

	service, db := newItemService(ctx, t)
	defer func() {
		_ = db.Close()
	}()

	mux := newMux(t, "POST /items", handlers.NewItemCreateHandler(service))

	t.Run("responds 400 on parse form fails", func(t *testing.T) {
		request := newInvalidPostRequest(ctx, t, "/items")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("renders errors on invalid input", func(t *testing.T) {
		cleanup(ctx, t)

		request := newPostRequest(ctx, t, "/items", Params{"date": ""})

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})

	t.Run("responds 200 on succcessful create", func(t *testing.T) {
		cleanup(ctx, t)

		categoryToCreate := &models.Category{
			ID:       1101,
			Name:     "Accoutrements",
			Currency: currencies.USD,
		}

		createCategory(ctx, t, categoryToCreate)

		params := Params{
			"date":        "2024-10-16",
			"formula":     "42.69+69.42",
			"category_id": "1101",
			"description": "paper clips, notebooks, and pens",
		}

		request := newPostRequest(ctx, t, "/items?currency=usd", params)

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)

		item := findItemByDate(ctx, t, currencies.USD, "2024-10-16")

		assert.Equal(t, item.GetDateAsString(), "2024-10-16")
		assert.Equal(t, item.CategoryID, 1101)
		assert.Equal(t, item.CategoryName, "Accoutrements")
		assert.Equal(t, item.Currency, currencies.USD)
		assert.Equal(t, item.Formula, "42.69+69.42")
		assert.Equal(t, item.GetSumAsString(), "112,11")
		assert.Equal(t, item.Description, "paper clips, notebooks, and pens")
	})
}
