package handlers_test

import (
	"database/sql"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/tksasha/balance/internal/backoffice/category/component"
	"github.com/tksasha/balance/internal/backoffice/category/handlers"
	"github.com/tksasha/balance/internal/backoffice/category/repository"
	"github.com/tksasha/balance/internal/backoffice/category/service"
	"github.com/tksasha/balance/internal/common"
	commonComponent "github.com/tksasha/balance/internal/common/component"
	"github.com/tksasha/balance/internal/common/currency"
	"github.com/tksasha/balance/internal/common/tests"
	"github.com/tksasha/balance/internal/db"
	nameprovider "github.com/tksasha/balance/internal/db/nameprovider/test"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCategoryCreateHandler(t *testing.T) { //nolint:funlen
	ctx := t.Context()

	handler, db := newCreateHandler(t)
	defer func() {
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	mux := tests.NewMux(t, "POST /backoffice/categories", handler)

	t.Run("responds 400 when input data is invalid", func(t *testing.T) {
		body := strings.NewReader("%")

		request, err := http.NewRequestWithContext(ctx, http.MethodPost, "/backoffice/categories", body)
		if err != nil {
			t.Fatal(err)
		}

		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusBadRequest)
	})

	t.Run("renders errors when validation failed", func(t *testing.T) {
		formData := url.Values{"name": {""}}

		body := strings.NewReader(formData.Encode())

		request, err := http.NewRequestWithContext(ctx, http.MethodPost, "/backoffice/categories", body)
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		responseBody := tests.GetResponseBody(t, recorder.Body)

		assert.Equal(t, recorder.Code, http.StatusOK)
		assert.Assert(t, is.Contains(responseBody, "name: is required"))
	})

	t.Run("responds 201 on successful create", func(t *testing.T) {
		tests.Cleanup2(ctx, t, db)

		formData := url.Values{
			"name":          {"Miscellaneous"},
			"income":        {"true"},
			"visible":       {"true"},
			"supercategory": {"3"},
		}

		t.Logf("%s", formData.Encode())

		body := strings.NewReader(formData.Encode())

		request, err := http.NewRequestWithContext(ctx, http.MethodPost, "/backoffice/categories?currency=eur", body)
		if err != nil {
			t.Fatal(err)
		}

		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		responseBody, err := io.ReadAll(recorder.Body)
		if err != nil {
			t.Fatal(err)
		}

		t.Logf("%s", responseBody)

		assert.Equal(t, recorder.Code, http.StatusCreated)

		category := tests.FindCategoryByName(ctx, t, currency.EUR, "Miscellaneous")

		assert.Equal(t, category.ID, 1)
		assert.Equal(t, category.Name, "Miscellaneous")
		assert.Equal(t, category.Income, true)
		assert.Equal(t, category.Visible, true)
		assert.Equal(t, category.Currency, currency.EUR)
		assert.Equal(t, category.Supercategory, 3)
	})
}

func newCreateHandler(t *testing.T) (*handlers.CreateHandler, *sql.DB) {
	t.Helper()

	db := db.Open(t.Context(), nameprovider.New())

	categoryRepository := repository.New(common.NewBaseRepository(), db)

	categoryService := service.New(common.NewBaseService(), categoryRepository)

	categoryComponent := component.New(commonComponent.New())

	handler := handlers.NewCreateHandler(common.NewBaseHandler(), categoryService, categoryComponent)

	return handler, db
}
