package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/backoffice/cash/components"
	"github.com/tksasha/balance/internal/backoffice/cash/handlers"
	"github.com/tksasha/balance/internal/common/tests"
	"gotest.tools/v3/assert"
)

func TestCashNewHandler(t *testing.T) {
	mux := tests.NewMux(t, "GET /cashes/new", newNewHandler(t))

	ctx := t.Context()

	t.Run("responds 200 when there are no errors", func(t *testing.T) {
		request, err := http.NewRequestWithContext(ctx, http.MethodGet, "/cashes/new", nil)
		if err != nil {
			t.Fatal(err)
		}

		responseWriter := httptest.NewRecorder()

		mux.ServeHTTP(responseWriter, request)

		assert.Equal(t, responseWriter.Code, http.StatusOK)
	})
}

func newNewHandler(t *testing.T) *handlers.NewHandler {
	t.Helper()

	cashComponent := components.NewCashComponent()

	handler := handlers.NewNewHandler(cashComponent)

	return handler
}
