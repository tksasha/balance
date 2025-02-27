package handlers_test

import (
	"database/sql"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/tksasha/balance/internal/app/cash"
	"github.com/tksasha/balance/internal/app/cash/handlers"
	"github.com/tksasha/balance/internal/app/cash/repository"
	"github.com/tksasha/balance/internal/app/cash/service"
	"github.com/tksasha/balance/internal/common/currency"
	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/db/nameprovider"
	"gotest.tools/v3/assert"
	"gotest.tools/v3/golden"
)

func TestCashUpdateHandler(t *testing.T) { //nolint:funlen
	handler, db := newUpdateHandler(t)
	defer func() {
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	mux := mux(t, "PATCH /cashes/{id}", handler)

	ctx := t.Context()

	t.Run("renders 404 when cash not found", func(t *testing.T) {
		values := url.Values{}

		body := strings.NewReader(values.Encode())

		request, err := http.NewRequestWithContext(ctx, http.MethodPatch, "/cashes/1439", body)
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNotFound)
	})

	t.Run("renders 400 when invalid input", func(t *testing.T) {
		request, err := http.NewRequestWithContext(ctx, http.MethodPatch, "/cashes/1453", nil)
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

		values := url.Values{}

		body := strings.NewReader(values.Encode())

		request, err := http.NewRequestWithContext(ctx, http.MethodPatch, "/cashes/1418?currency=usd", body)
		if err != nil {
			t.Fatal(err)
		}

		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)

		response, err := io.ReadAll(recorder.Body)
		if err != nil {
			t.Fatal(err)
		}

		golden.Assert(t, string(response), "edit-with-errors.html")
	})

	t.Run("renders updated cash when it updates", func(t *testing.T) {
		cleanup(t, db)

		cashToCreate := &cash.Cash{
			ID:            1442,
			Currency:      currency.UAH,
			Formula:       "2+3",
			Sum:           5,
			Name:          "Bonds",
			Supercategory: 2,
		}

		createCash(t, db, cashToCreate)

		values := url.Values{
			"formula": {"3+4"},
			"name":    {"Stocks"},
		}

		body := strings.NewReader(values.Encode())

		request, err := http.NewRequestWithContext(ctx, http.MethodPatch, "/cashes/1442", body)
		if err != nil {
			t.Fatal(err)
		}

		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)

		response, err := io.ReadAll(recorder.Body)
		if err != nil {
			t.Fatal(err)
		}

		golden.Assert(t, string(response), "update.html")

		cash := findCashByID(t, db, currency.UAH, 1442)

		assert.Equal(t, cash.ID, 1442)
		assert.Equal(t, cash.Currency, currency.UAH)
		assert.Equal(t, cash.Formula, "3+4")
		assert.Equal(t, cash.Sum, 7.0)
		assert.Equal(t, cash.Supercategory, 2)
	})
}

func newUpdateHandler(t *testing.T) (*handlers.UpdateHandler, *sql.DB) {
	t.Helper()

	db := db.Open(t.Context(), nameprovider.NewTestProvider())

	cashRepository := repository.New(db)

	cashService := service.New(cashRepository)

	handler := handlers.NewUpdateHandler(cashService)

	return handler, db
}
