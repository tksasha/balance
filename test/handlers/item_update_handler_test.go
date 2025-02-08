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

func TestItemUpdateHandler(t *testing.T) { //nolint:funlen
	ctx := context.Background()

	service, db := newItemService(ctx, t)
	defer func() {
		_ = db.Close()
	}()

	mux := newMux(t, "PATCH /items/{id}", handlers.NewItemUpdateHandler(service))

	t.Run("responds 400 on invalid input", func(t *testing.T) {
		cleanup(ctx, t)

		request := newInvalidPatchRequest(ctx, t, "/items/1138")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusBadRequest)
	})

	t.Run("responds 404 on no item found", func(t *testing.T) {
		cleanup(ctx, t)

		request := newPatchRequest(ctx, t, "/items/1218", Params{"date": "2025-01-25"})

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNotFound)
	})

	t.Run("responds 200 on successful update", func(t *testing.T) {
		cleanup(ctx, t)

		categoryToCreate := &models.Category{
			ID:       1148,
			Name:     "Pharmaceutical",
			Currency: currencies.EUR,
		}

		createCategory(ctx, t, categoryToCreate)

		itemToCreate := &models.Item{
			ID:         1143,
			CategoryID: 1148,
			Currency:   currencies.EUR,
		}

		createItem(ctx, t, itemToCreate)

		request := newPatchRequest(ctx, t, "/items/1143?currency=eur",
			Params{
				"date":        "2025-01-25",
				"formula":     "24 + 11 + 49",
				"category_id": "1148",
				"description": "pizza, ninja and disco",
			},
		)

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)

		item := findItemByDate(ctx, t, currencies.EUR, "2025-01-25")

		assert.Equal(t, item.ID, 1143)
	})
}
