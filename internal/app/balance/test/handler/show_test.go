package handler_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/app/balance/handler"
	"github.com/tksasha/balance/internal/app/balance/repository"
	"github.com/tksasha/balance/internal/app/balance/service"
	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/db/nameprovider"
	"github.com/tksasha/balance/internal/server/middlewares"
	"gotest.tools/v3/assert"
	"gotest.tools/v3/golden"
)

func TestShowHandler(t *testing.T) {
	ctx := t.Context()

	db := db.Open(ctx, nameprovider.NewTestProvider())

	balanceRepository := repository.New(db)

	balanceService := service.New(balanceRepository)

	showBalanceHandler := handler.NewShowHandler(balanceService)

	next := http.Handler(showBalanceHandler)

	for _, middleware := range middlewares.New() {
		next = middleware.Wrap(next)
	}

	mux := http.NewServeMux()

	mux.Handle("GET /balance", next)

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, "/balance", nil)
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

	golden.Assert(t, string(response), "balance.html")
}
