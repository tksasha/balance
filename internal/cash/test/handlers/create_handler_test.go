package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/tksasha/balance/internal/cash/handlers"
	"github.com/tksasha/balance/internal/common/testutils"
	"github.com/tksasha/balance/pkg/currencies"
	"github.com/tksasha/balance/test/handlers/utils"
	"gotest.tools/v3/assert"
)

func TestCashCreateHandler(t *testing.T) { //nolint:funlen
	ctx := t.Context()

	service, db := testutils.NewCashService(ctx, t)
	defer func() {
		_ = db.Close()
	}()

	handler := handlers.NewCreateHandler(service)

	mux := utils.NewMux(t, "POST /cashes", handler)

	t.Run("responds 400 whe parse form failed", func(t *testing.T) {
		request := utils.NewInvalidPostRequest(ctx, t, "/cashes")

		responseWriter := httptest.NewRecorder()

		mux.ServeHTTP(responseWriter, request)

		assert.Equal(t, responseWriter.Code, http.StatusBadRequest)
	})

	t.Run("responds 200 when input is invalid", func(t *testing.T) {
		request := utils.NewPostRequest(ctx, t, "/cashes", utils.Params{"name": ""})

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		body := utils.GetResponseBody(t, recorder.Body)

		assert.Equal(t, recorder.Code, http.StatusOK)
		assert.Assert(t, strings.Contains(body, "name: is required"))
	})

	t.Run("responds 200 when create succeeded", func(t *testing.T) {
		utils.Cleanup(ctx, t)

		params := utils.Params{
			"name":          "Bonds",
			"formula":       "2+3",
			"supercategory": "2",
			"favorite":      "true",
		}

		request := utils.NewPostRequest(ctx, t, "/cashes?currency=usd", params)

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)

		cash := utils.FindCashByName(ctx, t, currencies.USD, "Bonds")

		assert.Equal(t, cash.ID, 1)
		assert.Equal(t, cash.Name, "Bonds")
		assert.Equal(t, cash.Formula, "2+3")
		assert.Equal(t, cash.Sum, 5.0)
		assert.Equal(t, cash.Currency, currencies.USD)
		assert.Equal(t, cash.Supercategory, 2)
		assert.Equal(t, cash.Favorite, true)
	})
}
