package handlers_test

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/core/cash"
	"github.com/tksasha/balance/internal/core/cash/handlers"
	"github.com/tksasha/balance/internal/core/common/tests"
	"github.com/tksasha/balance/pkg/currencies"
	"gotest.tools/v3/assert"
)

func TestCashDeleteHandler(t *testing.T) {
	ctx := t.Context()

	service, db := tests.NewCashService(ctx, t)
	defer func() {
		_ = db.Close()
	}()

	handler := handlers.NewDeleteHandler(service)

	mux := tests.NewMux(t, "DELETE /cashes/{id}", handler)

	t.Run("renders 404 when cash not found", func(t *testing.T) {
		request := tests.NewDeleteRequest(ctx, t, "/cashes/1007")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNotFound)
	})

	t.Run("renders 204 when cash deleted", func(t *testing.T) {
		tests.Cleanup(ctx, t)

		cashToCreate := &cash.Cash{
			ID:        1011,
			Currency:  currencies.UAH,
			DeletedAt: sql.NullTime{},
		}

		tests.CreateCash(ctx, t, cashToCreate)

		request := tests.NewDeleteRequest(ctx, t, "/cashes/1011")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNoContent)

		cash := tests.FindCashByID(ctx, t, currencies.UAH, 1011)

		assert.Equal(t, cash.ID, 1011)
		assert.Assert(t, cash.DeletedAt.Valid)
	})
}
