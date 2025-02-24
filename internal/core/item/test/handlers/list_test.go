package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/common/tests"
	"gotest.tools/v3/assert"
)

func TestItemListHandler(t *testing.T) {
	ctx := t.Context()

	itemService, db := tests.NewItemService(ctx, t)
	defer func() {
		_ = db.Close()
	}()

	mux := tests.NewMux(t, "GET /items", tests.NewListItemsHandler(t, itemService))

	t.Run("responds 200 on items found", func(t *testing.T) {
		request := tests.NewGetRequest(ctx, t, "/items?currency=eur")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}
