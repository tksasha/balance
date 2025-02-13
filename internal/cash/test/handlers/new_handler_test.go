package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/cash/handlers"
	"github.com/tksasha/balance/internal/common/testutils"
	"github.com/tksasha/balance/test/handlers/utils"
	"gotest.tools/v3/assert"
)

func TestCashNewHandler(t *testing.T) {
	ctx := t.Context()

	mux := testutils.NewMux(t, "GET /cashes/new", handlers.NewNewHandler())

	t.Run("responds 200 when there are no errors", func(t *testing.T) {
		request := utils.NewGetRequest(ctx, t, "/cashes/new")

		responseWriter := httptest.NewRecorder()

		mux.ServeHTTP(responseWriter, request)

		assert.Equal(t, responseWriter.Code, http.StatusOK)
	})
}
