package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/tksasha/balance/internal/core/cash"
	"github.com/tksasha/balance/internal/core/cash/handlers"
	"github.com/tksasha/balance/internal/core/common/tests"
	"github.com/tksasha/balance/pkg/currencies"
	"gotest.tools/v3/assert"
)

func TestCashEditHandler(t *testing.T) { //nolint:funlen
	ctx := t.Context()

	service, db := tests.NewCashService(ctx, t)
	defer func() {
		_ = db.Close()
	}()

	handler := handlers.NewEditHandler(service)

	mux := tests.NewMux(t, "GET /cashes/{id}/edit", handler)

	t.Run("renders 404 on invalid id", func(t *testing.T) {
		request := tests.NewGetRequest(ctx, t, "/cashes/abc/edit")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNotFound)
	})

	t.Run("renders 404 on not found", func(t *testing.T) {
		tests.Cleanup(ctx, t)

		request := tests.NewGetRequest(ctx, t, "/cashes/1255/edit")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNotFound)
	})

	t.Run("renders form on success", func(t *testing.T) {
		tests.Cleanup(ctx, t)

		cash := &cash.Cash{
			ID:            1300,
			Currency:      currencies.EUR,
			Formula:       "2+3",
			Sum:           5,
			Name:          "Bonds",
			Supercategory: 6,
			Favorite:      true,
		}

		tests.CreateCash(ctx, t, cash)

		request := tests.NewGetRequest(ctx, t, "/cashes/1300/edit?currency=eur")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)

		body := tests.GetResponseBody(t, recorder.Body)

		assert.Assert(t, strings.Contains(body, "Bonds"))
		assert.Assert(t, strings.Contains(body, "2+3"))
	})
}
