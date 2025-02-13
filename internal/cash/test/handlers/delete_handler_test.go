package handlers_test

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/cash"
	"github.com/tksasha/balance/internal/cash/handlers"
	"github.com/tksasha/balance/internal/common/testutils"
	"github.com/tksasha/balance/pkg/currencies"
	"gotest.tools/v3/assert"
)

func TestCashDeleteHandler(t *testing.T) {
	ctx := t.Context()

	service, db := testutils.NewCashService(ctx, t)
	defer func() {
		_ = db.Close()
	}()

	handler := handlers.NewDeleteHandler(service)

	mux := testutils.NewMux(t, "DELETE /cashes/{id}", handler)

	t.Run("renders 404 when cash not found", func(t *testing.T) {
		request := testutils.NewDeleteRequest(ctx, t, "/cashes/1007")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNotFound)
	})

	t.Run("renders 204 when cash deleted", func(t *testing.T) {
		testutils.Cleanup(ctx, t)

		cashToCreate := &cash.Cash{
			ID:        1011,
			Currency:  currencies.UAH,
			DeletedAt: sql.NullTime{},
		}

		testutils.CreateCash(ctx, t, cashToCreate)

		request := testutils.NewDeleteRequest(ctx, t, "/cashes/1011")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNoContent)

		cash := testutils.FindCashByID(ctx, t, currencies.UAH, 1011)

		assert.Equal(t, cash.ID, 1011)
		assert.Assert(t, cash.DeletedAt.Valid)
	})
}
