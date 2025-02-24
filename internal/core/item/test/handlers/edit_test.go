package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/common/currency"
	"github.com/tksasha/balance/internal/common/tests"
	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/core/item"
	"gotest.tools/v3/assert"
)

func TestItemEditHandler(t *testing.T) {
	ctx := t.Context()

	itemService, db := tests.NewItemService(ctx, t)
	defer func() {
		_ = db.Close()
	}()

	categoryService, db2 := tests.NewCategoryService(ctx, t)
	defer func() {
		_ = db2.Close()
	}()

	mux := tests.NewMux(t, "GET /items/{id}/edit", tests.NewEditItemHandler(t, itemService, categoryService))

	t.Run("responds 404 when item not found", func(t *testing.T) {
		request := tests.NewGetRequest(ctx, t, "/items/1514/edit?currency=usd")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusNotFound, recorder.Code)
	})

	t.Run("responds 200 when item found", func(t *testing.T) {
		tests.Cleanup(ctx, t)

		categoryToCreate := &category.Category{
			ID:       5,
			Currency: currency.UAH,
		}

		tests.CreateCategory(ctx, t, categoryToCreate)

		itemToCreate := &item.Item{
			ID:         1745,
			Currency:   currency.UAH,
			CategoryID: 5,
		}

		tests.CreateItem(ctx, t, itemToCreate)

		request := tests.NewGetRequest(ctx, t, "/items/1745/edit")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)
	})
}
