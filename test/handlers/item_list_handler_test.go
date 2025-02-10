package handlers_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/handlers"
	"gotest.tools/v3/assert"
)

func TestItemListHandler(t *testing.T) {
	ctx := context.Background()

	service, db := newItemService(ctx, t)
	defer func() {
		_ = db.Close()
	}()

	mux := newMux(t, "GET /items", handlers.NewItemListHandler(service))

	t.Run("responds 200 on items found", func(t *testing.T) {
		request := newGetRequest(ctx, t, "/items?currency=eur")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}
