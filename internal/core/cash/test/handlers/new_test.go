package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/core/common/tests"
	"gotest.tools/v3/assert"
)

func TestCashNewHandler(t *testing.T) {
	ctx := t.Context()

	mux := tests.NewMux(t, "GET /cashes/new", tests.NewNewCasheHandler(t))

	t.Run("responds 200 when there are no errors", func(t *testing.T) {
		request := tests.NewGetRequest(ctx, t, "/cashes/new")

		responseWriter := httptest.NewRecorder()

		mux.ServeHTTP(responseWriter, request)

		assert.Equal(t, responseWriter.Code, http.StatusOK)
	})
}
