package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/tksasha/balance/internal/backoffice/cash"
	"github.com/tksasha/balance/internal/backoffice/cash/handlers"
	"github.com/tksasha/balance/internal/backoffice/cash/repository"
	"github.com/tksasha/balance/internal/backoffice/cash/service"
	"github.com/tksasha/balance/internal/common/currency"
	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/db/nameprovider"
	"github.com/tksasha/balance/internal/server/middlewares"
	"gotest.tools/v3/assert"
)

func TestCashUpdateHandler(t *testing.T) { //nolint:funlen
	ctx := t.Context()

	db := db.Open(ctx, nameprovider.NewTestProvider())
	defer func() {
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	cashRepository := repository.New(db)

	cashService := service.New(cashRepository)

	cashUpdateHandler := handlers.NewUpdateHandler(cashService)

	mux := http.NewServeMux()

	next := http.Handler(cashUpdateHandler)

	for _, middleware := range middlewares.New() {
		next = middleware.Wrap(next)
	}

	mux.Handle("PATCH /backoffice/cashes/{id}", next)

	t.Run("renders 404 when cash not found", func(t *testing.T) {
		formData := url.Values{}

		body := strings.NewReader(formData.Encode())

		request, err := http.NewRequestWithContext(ctx, http.MethodPatch, "/backoffice/cashes/1439", body)
		if err != nil {
			t.Fatal(err)
		}

		request.Header.Add("Content-Type", "application/x-form-www-urlencoded")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNotFound)
	})

	t.Run("renders 400 when invalid input", func(t *testing.T) {
		request, err := http.NewRequestWithContext(ctx, http.MethodPatch, "/backoffice/cashes/1453", nil)
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusBadRequest)
	})

	t.Run("renders errors when validation failed", func(t *testing.T) {
		cleanup(t, db)

		cashToCreate := &cash.Cash{
			ID:       1418,
			Currency: currency.USD,
		}

		createCash(t, db, cashToCreate)

		values := url.Values{
			"name": {""},
		}

		body := strings.NewReader(values.Encode())

		request, err := http.NewRequestWithContext(ctx, http.MethodPatch, "/backoffice/cashes/1418?currency=usd", body)
		if err != nil {
			t.Fatal(err)
		}

		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)
	})

	t.Run("renders 200 when cash updated", func(t *testing.T) {
		cleanup(t, db)

		cashToCreate := &cash.Cash{
			ID:            1442,
			Currency:      currency.UAH,
			Formula:       "2+3",
			Sum:           5,
			Name:          "Bonds",
			Supercategory: 2,
			Favorite:      false,
		}

		createCash(t, db, cashToCreate)

		values := url.Values{
			"formula":       {"3+4"},
			"name":          {"Stocks"},
			"supercategory": {"3"},
			"favorite":      {"true"},
		}

		body := strings.NewReader(values.Encode())

		request, err := http.NewRequestWithContext(ctx, http.MethodPatch, "/backoffice/cashes/1442", body)
		if err != nil {
			t.Fatal(err)
		}

		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)

		expectedHeader := `{"backoffice.cash.updated":{"backofficeCashesPath":"/backoffice/cashes"}}`

		assert.Equal(t, expectedHeader, strings.TrimSpace(recorder.Header().Get("Hx-Trigger-After-Swap")))
	})
}
