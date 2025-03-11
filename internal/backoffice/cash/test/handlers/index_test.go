package handlers_test

import (
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
)

func TestCashListHandler(t *testing.T) {
	db := db.Open(t.Context(), nameprovider.NewTestProvider())
	defer func() {
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	cashRepository := repository.New(db)

	cashService := service.New(cashRepository)

	handler := handlers.NewIndexHandler(cashService)

	mux := http.NewServeMux()

	mux.Handle("GET /backoffice/cashes", handler)

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

		assert.Equal(t, "backoffice.cashes.shown", recorder.Header().Get("Hx-Trigger-After-Swap"))
	})
}
