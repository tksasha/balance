package handlers_test

import (
	"context"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/handlers"
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/pkg/currencies"
	"gotest.tools/v3/assert"
)

func TestCashDeleteHandler(t *testing.T) {
	ctx := context.Background()

	service, db := newCashService(ctx, t)
	defer func() {
		_ = db.Close()
	}()

	handler := handlers.NewCashDeleteHandler(service)

	mux := newMux(t, "DELETE /cashes/{id}", handler)

	_ = mux

	t.Run("renders 404 when cash not found", func(t *testing.T) {
		request := newDeleteRequest(ctx, t, "/cashes/1007")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNotFound)
	})

	t.Run("renders 200 when cash found", func(t *testing.T) {
		cleanup(ctx, t)

		cashToCreate := &models.Cash{
			ID:        1011,
			Currency:  currencies.UAH,
			DeletedAt: sql.NullTime{},
		}

		createCash(ctx, t, cashToCreate)

		request := newDeleteRequest(ctx, t, "/cashes/1011")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)

		cash := findCashByID(ctx, t, currencies.UAH, 1011)

		assert.Equal(t, cash.ID, 1011)
		assert.Assert(t, cash.DeletedAt.Valid)
	})
}
