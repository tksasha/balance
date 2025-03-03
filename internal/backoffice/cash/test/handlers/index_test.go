package handlers_test

import (
	"database/sql"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/backoffice/cash"
	"github.com/tksasha/balance/internal/backoffice/cash/handlers"
	"github.com/tksasha/balance/internal/backoffice/cash/repository"
	"github.com/tksasha/balance/internal/backoffice/cash/service"
	"github.com/tksasha/balance/internal/common/currency"
	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/db/nameprovider"
	"gotest.tools/v3/assert"
	"gotest.tools/v3/golden"
)

func TestCashListHandler(t *testing.T) {
	handler, db := newListHandler(t)
	defer func() {
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	mux := mux(t, "GET /backoffice/cashes", handler)

	ctx := t.Context()

	t.Run("renders cash list when there no errors", func(t *testing.T) {
		cleanup(t, db)

		testdata := map[int]string{1: "Stocks", 2: "Bonds", 3: "Cash"}

		for id, name := range testdata {
			cash := &cash.Cash{
				ID:       id,
				Name:     name,
				Currency: currency.Default,
			}

			createCash(t, db, cash)
		}

		request, err := http.NewRequestWithContext(ctx, http.MethodGet, "/backoffice/cashes", nil)
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

func newListHandler(t *testing.T) (*handlers.ListHandler, *sql.DB) {
	t.Helper()

	db := db.Open(t.Context(), nameprovider.NewTestProvider())

	cashRepository := repository.New(db)

	cashService := service.New(cashRepository)

	handler := handlers.NewListHandler(cashService)

	return handler, db
}
