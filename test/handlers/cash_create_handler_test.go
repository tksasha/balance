package handlers_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/handlers"
	"github.com/tksasha/balance/internal/middlewares"
	providers "github.com/tksasha/balance/internal/providers/test"
	"github.com/tksasha/balance/internal/repositories"
	"github.com/tksasha/balance/internal/services"
	"github.com/tksasha/balance/pkg/currencies"
	"gotest.tools/v3/assert"
)

func TestCashCreateHandler(t *testing.T) { //nolint:funlen
	dbNameProvider := providers.NewDBNameProvider()

	db := db.Open(dbNameProvider)

	cashRepository := repositories.NewCashRepository(db)

	cashService := services.NewCashService(cashRepository)

	middleware := middlewares.NewCurrencyMiddleware().Wrap(
		handlers.NewCashCreateHandler(cashService),
	)

	mux := http.NewServeMux()
	mux.Handle("POST /cashes", middleware)

	ctx := context.Background()

	t.Run("responds 400 whe parse form failed", func(t *testing.T) {
		request := newInvalidPostRequest(ctx, t, "/cashes")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusBadRequest)
	})

	t.Run("responds 200 when input is invalid", func(t *testing.T) {
		request := newPostRequest(ctx, t, "/cashes", Params{"name": ""})

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		body := getResponseBody(t, recorder.Body)

		assert.Equal(t, recorder.Code, http.StatusOK)
		assert.Assert(t, strings.Contains(body, "name: is required"))
	})

	t.Run("responds 200 when create succeeded", func(t *testing.T) {
		cleanup(ctx, t, db)

		params := Params{
			"name":          "Bonds",
			"formula":       "2+3",
			"supercategory": "2",
			"favorite":      "true",
		}

		request := newPostRequest(ctx, t, "/cashes?currency=usd", params)

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)

		cash := findCashByName(usdContext(ctx, t), t, db, "Bonds")

		assert.Equal(t, cash.ID, 1)
		assert.Equal(t, cash.Name, "Bonds")
		assert.Equal(t, cash.Formula, "2+3")
		assert.Equal(t, cash.Sum, 5.0)
		assert.Equal(t, cash.Currency, currencies.USD)
		assert.Equal(t, cash.Supercategory, 2)
		assert.Equal(t, cash.Favorite, true)
	})
}
