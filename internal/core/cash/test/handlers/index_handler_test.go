package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/core/cash/handlers"
	"github.com/tksasha/balance/internal/core/common/tests"
	"gotest.tools/v3/assert"
)

func TestCashIndexHandler(t *testing.T) {
	ctx := t.Context()

	cashService, db := tests.NewCashService(ctx, t)
	defer func() {
		if err := db.Close(); err != nil {
			t.Logf("failed to close db: %v", err)
		}
	}()

	mux := tests.NewMux(t, "GET /cashes", handlers.NewIndexHandler(cashService))

	t.Run("renders cash list when there no errors", func(t *testing.T) {
		request := tests.NewGetRequest(ctx, t, "/cashes")

		responseWriter := httptest.NewRecorder()

		mux.ServeHTTP(responseWriter, request)

		assert.Equal(t, responseWriter.Code, http.StatusOK)
	})
}
