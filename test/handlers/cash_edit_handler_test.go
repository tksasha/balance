package handlers_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/handlers"
	"github.com/tksasha/balance/internal/models"
	providers "github.com/tksasha/balance/internal/providers/test"
	"github.com/tksasha/balance/internal/repositories"
	"github.com/tksasha/balance/internal/services"
	"github.com/tksasha/balance/pkg/currencies"
	"gotest.tools/v3/assert"
)

func TestCashEditHandler(t *testing.T) { //nolint:funlen
	dbNameProvider := providers.NewDBNameProvider()

	db := db.Open(dbNameProvider)

	cashRepository := repositories.NewCashRepository(db)

	cashService := services.NewCashService(cashRepository)

	mux := newMux(t, "GET /cashes/{id}/edit", handlers.NewCashEditHandler(cashService))

	ctx := context.Background()

	t.Run("renders 404 on invalid id", func(t *testing.T) {
		request := newGetRequest(ctx, t, "/cashes/abc/edit")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNotFound)
	})

	t.Run("renders 404 on not found", func(t *testing.T) {
		cleanup(ctx, t, db)

		request := newGetRequest(ctx, t, "/cashes/1255/edit")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNotFound)
	})

	t.Run("renders resource on success", func(t *testing.T) {
		cleanup(ctx, t, db)

		cash := &models.Cash{
			ID:            1300,
			Currency:      currencies.EUR,
			Formula:       "2+3",
			Sum:           5,
			Name:          "Bonds",
			Supercategory: 6,
			Favorite:      true,
		}

		createCash(ctx, t, db, cash)

		request := newGetRequest(ctx, t, "/cashes/1300/edit?currency=eur")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)

		body := getResponseBody(t, recorder.Body)

		assert.Assert(t, strings.Contains(body, "Bonds"))
	})
}
