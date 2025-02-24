package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/common/currency"
	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/common/tests"
	"github.com/tksasha/balance/internal/core/item"
	"gotest.tools/v3/assert"
)

func TestItemUpdateHandler(t *testing.T) { //nolint:funlen
	ctx := t.Context()

	itemService, db := tests.NewItemService(ctx, t)
	defer func() {
		_ = db.Close()
	}()

	categoryService, db2 := tests.NewCategoryService(ctx, t)
	defer func() {
		_ = db2.Close()
	}()

	mux := tests.NewMux(t, "PATCH /items/{id}", tests.NewUpdateItemHandler(t, itemService, categoryService))

	t.Run("responds 400 on invalid input", func(t *testing.T) {
		tests.Cleanup(ctx, t)

		request := tests.NewInvalidPatchRequest(ctx, t, "/items/1138")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusBadRequest)
	})

	t.Run("responds 404 on no item found", func(t *testing.T) {
		tests.Cleanup(ctx, t)

		params := tests.Params{"date": "2025-01-25"}

		request := tests.NewPatchRequest(ctx, t, "/items/1218", params)

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNotFound)
	})

	t.Run("responds 204 on successful update", func(t *testing.T) {
		tests.Cleanup(ctx, t)

		categoryToCreate := &category.Category{
			ID:       1148,
			Name:     "Pharmaceutical",
			Currency: currency.EUR,
		}

		tests.CreateCategory(ctx, t, categoryToCreate)

		itemToCreate := &item.Item{
			ID:         1143,
			CategoryID: 1148,
			Currency:   currency.EUR,
		}

		tests.CreateItem(ctx, t, itemToCreate)

		params := tests.Params{
			"date":        "2025-01-25",
			"formula":     "24 + 11 + 49",
			"category_id": "1148",
			"description": "pizza, ninja and disco",
		}

		request := tests.NewPatchRequest(ctx, t, "/items/1143?currency=eur", params)

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNoContent)

		item := tests.FindItemByDate(ctx, t, currency.EUR, "2025-01-25")

		assert.Equal(t, item.ID, 1143)
	})
}
