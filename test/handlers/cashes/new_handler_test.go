package handlers_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/handlers/cashes"
	"github.com/tksasha/balance/test/handlers/utils"
	"gotest.tools/v3/assert"
)

func TestCashNewHandler(t *testing.T) {
	ctx := context.Background()

	mux := utils.NewMux(t, "GET /cashes/new", cashes.NewNewHandler())

	t.Run("responds 200 when there are no errors", func(t *testing.T) {
		request := utils.NewGetRequest(ctx, t, "/cashes/new")

		responseWriter := httptest.NewRecorder()

		mux.ServeHTTP(responseWriter, request)

		assert.Equal(t, responseWriter.Code, http.StatusOK)
	})
}
