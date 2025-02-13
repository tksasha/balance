package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/cash"
	"github.com/tksasha/balance/internal/cash/handlers"
	"github.com/tksasha/balance/internal/common/testutils"
	"github.com/tksasha/balance/pkg/currencies"
	"gotest.tools/v3/assert"
)

func TestCashUpdateHandler(t *testing.T) {
	ctx := t.Context()

	service, db := testutils.NewCashService(ctx, t)
	defer func() {
		_ = db.Close()
	}()

	mux := testutils.NewMux(t, "PATCH /cashes/{id}", handlers.NewUpdateHandler(service))

	t.Run("renders 404 when cash not found", func(t *testing.T) {
		request := testutils.NewPatchRequest(ctx, t, "/cashes/1439", nil)

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNotFound)
	})

	t.Run("renders 400 when invalid input", func(t *testing.T) {
		request := testutils.NewInvalidPatchRequest(ctx, t, "/cashes/1453")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusBadRequest)
	})

	t.Run("renders 200 when cash updated", func(t *testing.T) {
		testutils.Cleanup(ctx, t)

		cashToCreate := &cash.Cash{
			ID:            1442,
			Currency:      currencies.UAH,
			Formula:       "2+3",
			Sum:           5,
			Name:          "Bonds",
			Supercategory: 2,
			Favorite:      false,
		}

		testutils.CreateCash(ctx, t, cashToCreate)

		params := testutils.Params{
			"formula":       "3+4",
			"name":          "Stocks",
			"supercategory": "3",
			"favorite":      "true",
		}

		request := testutils.NewPatchRequest(ctx, t, "/cashes/1442", params)

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)
	})
}
