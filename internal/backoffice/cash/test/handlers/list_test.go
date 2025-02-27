package handlers_test

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/backoffice/cash/handlers"
	"github.com/tksasha/balance/internal/backoffice/cash/repository"
	"github.com/tksasha/balance/internal/backoffice/cash/service"
	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/db/nameprovider"
	"gotest.tools/v3/assert"
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
		request, err := http.NewRequestWithContext(ctx, http.MethodGet, "/backoffice/cashes", nil)
		if err != nil {
			t.Fatal(err)
		}

		responseWriter := httptest.NewRecorder()

		mux.ServeHTTP(responseWriter, request)

		assert.Equal(t, responseWriter.Code, http.StatusOK)
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
