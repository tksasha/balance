package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/core/common/helpers"
	"github.com/tksasha/balance/internal/core/common/providers"
	"github.com/tksasha/balance/internal/core/common/tests"
	"github.com/tksasha/balance/internal/core/item/handlers"
	"gotest.tools/v3/assert"
)

func TestItemIndexHandler(t *testing.T) {
	ctx := t.Context()

	service, db := tests.NewItemService(ctx, t)
	defer func() {
		_ = db.Close()
	}()

	timeProvider := providers.NewTimeProvider()

	helpers := helpers.New(timeProvider)

	mux := tests.NewMux(t, "GET /items", handlers.NewIndexHandler(service, helpers))

	t.Run("responds 200 on items found", func(t *testing.T) {
		request := tests.NewGetRequest(ctx, t, "/items?currency=eur")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}
