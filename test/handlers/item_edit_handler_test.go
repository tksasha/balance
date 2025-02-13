package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/handlers"
	"gotest.tools/v3/assert"
)

func TestItemEditHandler(t *testing.T) {
	ctx := t.Context()

	service, db := newItemService(ctx, t)
	defer func() {
		_ = db.Close()
	}()

	mux := newMux(t, "GET /items/{id}/edit", handlers.NewItemEditHandler(service))

	t.Run("responds 404 on no item found", func(t *testing.T) {
		request := newGetRequest(ctx, t, "/items/1514/edit?currency=usd")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusNotFound, recorder.Code)
	})
}
