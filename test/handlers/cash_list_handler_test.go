package handlers_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/handlers"
	"gotest.tools/v3/assert"
)

func TestCashListHandler(t *testing.T) {
	ctx := context.Background()

	cashService, db := newCashService(ctx, t)
	defer func() {
		if err := db.Close(); err != nil {
			t.Logf("failed to close db: %v", err)
		}
	}()

	mux := newMux(t, "GET /cashes", handlers.NewCashListHandler(cashService))

	t.Run("renders cash list when there no errors", func(t *testing.T) {
		request := newGetRequest(ctx, t, "/cashes")

		responseWriter := httptest.NewRecorder()

		mux.ServeHTTP(responseWriter, request)

		assert.Equal(t, responseWriter.Code, http.StatusOK)
	})
}
