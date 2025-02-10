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

func TestCashUpdateHandler(t *testing.T) {
	ctx := context.Background()

	service, db := newCashService(ctx, t)
	defer func() {
		_ = db.Close()
	}()

	mux := newMux(t, "PATCH /cashes/{id}", handlers.NewCashUpdateHandler(service))

	t.Run("renders 404 when cash not found", func(t *testing.T) {
		request := newPatchRequest(ctx, t, "/cashes/1439", nil)

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNotFound)
	})

	t.Run("renders 400 when invalid input", func(t *testing.T) {
		request := newInvalidPatchRequest(ctx, t, "/cashes/1453")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusBadRequest)
	})

	t.Run("renders 200 when cash updated", func(t *testing.T) {
		cleanup(ctx, t)

		cashToCreate := &models.Cash{
			ID:            1442,
			Currency:      currencies.UAH,
			Formula:       "2+3",
			Sum:           5,
			Name:          "Bonds",
			Supercategory: 2,
			Favorite:      false,
		}

		createCash(ctx, t, cashToCreate)

		params := Params{
			"formula":       "3+4",
			"name":          "Stocks",
			"supercategory": "3",
			"favorite":      "true",
		}

		request := newPatchRequest(ctx, t, "/cashes/1442", params)

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)
	})
}
