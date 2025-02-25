package handlers_test

import (
	"database/sql"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/tksasha/balance/internal/backoffice/cash/components"
	"github.com/tksasha/balance/internal/backoffice/cash/handlers"
	"github.com/tksasha/balance/internal/backoffice/cash/repository"
	"github.com/tksasha/balance/internal/backoffice/cash/service"
	"github.com/tksasha/balance/internal/common/currency"
	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/db/nameprovider"
	"gotest.tools/v3/assert"
)

func TestCashCreateHandler(t *testing.T) { //nolint:funlen
	handler, db := newCreateHandler(t)
	defer func() {
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	mux := mux(t, "POST /backoffice/cashes", handler)

	ctx := t.Context()

	t.Run("responds 400 whe parse form failed", func(t *testing.T) {
		request, err := http.NewRequestWithContext(ctx, http.MethodPost, "/backoffice/cashes", nil)
		if err != nil {
			t.Fatal(err)
		}

		responseWriter := httptest.NewRecorder()

		mux.ServeHTTP(responseWriter, request)

		assert.Equal(t, responseWriter.Code, http.StatusBadRequest)
	})

	t.Run("responds 200 when input is invalid", func(t *testing.T) {
		formData := url.Values{"name": {""}}

		request, err := http.NewRequestWithContext(
			ctx,
			http.MethodPost,
			"/backoffice/cashes",
			strings.NewReader(formData.Encode()),
		)
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		body, err := io.ReadAll(recorder.Body)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, recorder.Code, http.StatusOK)
		assert.Assert(t, strings.Contains(string(body), "name: is required"))
	})

	t.Run("responds 201 when create succeeded", func(t *testing.T) {
		cleanup(t, db)

		formData := url.Values{
			"name":          {"Bonds"},
			"formula":       {"2+3"},
			"supercategory": {"2"},
			"favorite":      {"true"},
		}

		request, err := http.NewRequestWithContext(
			ctx,
			http.MethodPost,
			"/backoffice/cashes?currency=usd",
			strings.NewReader(formData.Encode()),
		)
		if err != nil {
			t.Fatal(err)
		}

		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusCreated)

		cash := findCashByName(t, db, currency.USD, "Bonds")

		assert.Equal(t, cash.ID, 1)
		assert.Equal(t, cash.Name, "Bonds")
		assert.Equal(t, cash.Formula, "2+3")
		assert.Equal(t, cash.Sum, 5.0)
		assert.Equal(t, cash.Currency, currency.USD)
		assert.Equal(t, cash.Supercategory, 2)
		assert.Equal(t, cash.Favorite, true)
	})
}

func newCreateHandler(t *testing.T) (*handlers.CreateHandler, *sql.DB) {
	t.Helper()

	db := db.Open(t.Context(), nameprovider.NewTestProvider())

	cashRepository := repository.New(db)

	cashService := service.New(cashRepository)

	cashComponent := components.NewCashComponent()

	handler := handlers.NewCreateHandler(cashService, cashComponent)

	return handler, db
}
