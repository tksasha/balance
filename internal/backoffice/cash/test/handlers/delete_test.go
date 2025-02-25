package handlers_test

import (
	"database/sql"
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

func TestCashDeleteHandler(t *testing.T) {
	handler, db := newDeleteHandler(t)
	defer func() {
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	mux := mux(t, "DELETE /cashes/{id}", handler)

	ctx := t.Context()

	t.Run("renders 404 when cash not found", func(t *testing.T) {
		request, err := http.NewRequestWithContext(ctx, http.MethodDelete, "/cashes/1007", nil)
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNotFound)
	})

	t.Run("renders 204 when cash deleted", func(t *testing.T) {
		cleanup(t, db)

		cashToCreate := &cash.Cash{
			ID:        1011,
			Currency:  currency.UAH,
			DeletedAt: sql.NullTime{},
		}

		createCash(t, db, cashToCreate)

		request, err := http.NewRequestWithContext(ctx, http.MethodDelete, "/cashes/1011", nil)
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNoContent)

		cash := findCashByID(t, db, currency.UAH, 1011)

		assert.Equal(t, cash.ID, 1011)
		assert.Assert(t, cash.DeletedAt.Valid)
	})
}

func newDeleteHandler(t *testing.T) (*handlers.DeleteHandler, *sql.DB) {
	t.Helper()

	db := db.Open(t.Context(), nameprovider.NewTestProvider())

	cashRepository := repository.New(db)

	cashService := service.New(cashRepository)

	handler := handlers.NewDeleteHandler(cashService)

	return handler, db
}
