package handlers_test

import (
	"database/sql"
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
	"gotest.tools/v3/assert"
)

func TestCashEditHandler(t *testing.T) { //nolint:funlen
	handler, db := newEditHandler(t)
	defer func() {
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	mux := mux(t, "GET /cashes/{id}/edit", handler)

	ctx := t.Context()

	t.Run("renders 404 on invalid id", func(t *testing.T) {
		request, err := http.NewRequestWithContext(ctx, http.MethodGet, "/cashes/abc/edit", nil)
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNotFound)
	})

	t.Run("renders 404 on not found", func(t *testing.T) {
		request, err := http.NewRequestWithContext(ctx, http.MethodGet, "/cashes/1255/edit", nil)
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNotFound)
	})

	t.Run("renders form on success", func(t *testing.T) {
		cleanup(t, db)

		cashToCreate := &cash.Cash{
			ID:            1300,
			Currency:      currency.EUR,
			Formula:       "2+3",
			Sum:           5,
			Name:          "Bonds",
			Supercategory: 6,
		}

		createCash(t, db, cashToCreate)

		request, err := http.NewRequestWithContext(ctx, http.MethodGet, "/cashes/1300/edit?currency=eur", nil)
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)

		assert.Equal(t, "balance.cash.edit", recorder.Header().Get("Hx-Trigger-After-Swap"))
	})
}

func newEditHandler(t *testing.T) (*handlers.EditHandler, *sql.DB) {
	t.Helper()

	db := db.Open(t.Context(), nameprovider.NewTestProvider())

	cashRepository := repository.New(db)

	cashService := service.New(cashRepository)

	handler := handlers.NewEditHandler(cashService)

	return handler, db
}
