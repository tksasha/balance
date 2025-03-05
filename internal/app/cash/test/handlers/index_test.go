package handlers_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/app/cash"
	"github.com/tksasha/balance/internal/app/cash/handlers"
	"github.com/tksasha/balance/internal/app/cash/repository"
	"github.com/tksasha/balance/internal/app/cash/service"
	"github.com/tksasha/balance/internal/common/currency"
	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/db/nameprovider"
	"github.com/tksasha/balance/internal/server/middlewares"
	"gotest.tools/v3/assert"
	"gotest.tools/v3/golden"
)

func TestIndexHandler(t *testing.T) {
	ctx := t.Context()

	db := db.Open(ctx, nameprovider.NewTestProvider())
	defer func() {
		if err := db.Close(); err != nil {
			t.Logf("failed to close db: %v", err)
		}
	}()

	cashRepository := repository.New(db)

	cashService := service.New(cashRepository)

	indexHandler := handlers.NewIndexHandler(cashService)

	next := http.Handler(indexHandler)

	for _, middleware := range middlewares.New() {
		next = middleware.Wrap(next)
	}

	mux := http.NewServeMux()

	mux.Handle("/cashes", next)

	t.Run("render index.html", func(t *testing.T) {
		cleanup(t, db)

		createCash(t, db, &cash.Cash{ID: 1, Name: "Bonds", Sum: 16.45, Currency: currency.Default})

		request, err := http.NewRequestWithContext(ctx, http.MethodGet, "/cashes", nil)
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)

		response, err := io.ReadAll(recorder.Body)
		if err != nil {
			t.Fatal(err)
		}

		golden.Assert(t, string(response), "index.html")
	})
}
