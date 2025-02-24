package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/tksasha/balance/internal/common/currency"
	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/core/common/tests"
	"gotest.tools/v3/assert"
)

func TestItemCreateHandler(t *testing.T) { //nolint:funlen
	ctx := t.Context()

	itemService, db := tests.NewItemService(ctx, t)
	defer func() {
		_ = db.Close()
	}()

	categoryService, db2 := tests.NewCategoryService(ctx, t)
	defer func() {
		_ = db2.Close()
	}()

	mux := tests.NewMux(t, "POST /items", tests.NewCreateItemHandler(t, itemService, categoryService))

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

	t.Run("responds 204 when item created", func(t *testing.T) {
		tests.Cleanup(ctx, t)

		categoryToCreate := &category.Category{
			ID:       1101,
			Name:     "Accoutrements",
			Currency: currency.USD,
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

		assert.Equal(t, recorder.Code, http.StatusNoContent)

		item := tests.FindItemByDate(ctx, t, currency.USD, "2024-10-16")

		assert.Equal(t, item.Date.Format(time.DateOnly), "2024-10-16")
		assert.Equal(t, item.CategoryID, 1101)
		assert.Equal(t, item.CategoryName.String, "Accoutrements")
		assert.Equal(t, item.Currency, currency.USD)
		assert.Equal(t, item.Formula, "42.69+69.42")
		assert.Equal(t, item.Sum, 112.11)
		assert.Equal(t, item.Description, "paper clips, notebooks, and pens")
	})
}
