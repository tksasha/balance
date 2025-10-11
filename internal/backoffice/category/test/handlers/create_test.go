package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/tksasha/balance/internal/backoffice/category/handlers"
	"github.com/tksasha/balance/internal/backoffice/category/repository"
	"github.com/tksasha/balance/internal/backoffice/category/service"
	"github.com/tksasha/balance/internal/common/currency"
	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/db/nameprovider"
	"gotest.tools/v3/assert"
	"gotest.tools/v3/golden"
)

func TestCategoryCreateHandler(t *testing.T) { //nolint:funlen
	db := db.Open(t.Context(), nameprovider.NewTestProvider())

	defer func() {
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	categoryRepository := repository.New(db)

	categoryService := service.New(categoryRepository)

	handler := handlers.NewCreateHandler(categoryService)

	mux := mux(t, "POST /backoffice/categories", handler)

	ctx := t.Context()

	t.Run("responds 400 when input data is invalid", func(t *testing.T) {
		request, err := http.NewRequestWithContext(ctx, http.MethodPost, "/backoffice/categories", nil)
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

		assert.Equal(t, recorder.Code, http.StatusOK)
	})

	t.Run("responds 200 on successful create", func(t *testing.T) {
		cleanup(t, db)

		formData := url.Values{
			"name":          {"Miscellaneous"},
			"income":        {"true"},
			"visible":       {"true"},
			"supercategory": {"3"},
		}

		body := strings.NewReader(formData.Encode())

		request, err := http.NewRequestWithContext(ctx, http.MethodPost, "/backoffice/categories?currency=eur", body)
		if err != nil {
			t.Fatal(err)
		}

		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)

		golden.Assert(t, recorder.Header().Get("Hx-Trigger-After-Swap"),
			"create-hx-trigger-after-swap-header.json")

		category := findCategoryByName(t, db, currency.EUR, "Miscellaneous")

		assert.Equal(t, category.Name, "Miscellaneous")
		assert.Equal(t, category.Slug, "miscellaneous")
		assert.Equal(t, category.Income, true)
		assert.Equal(t, category.Visible, true)
		assert.Equal(t, category.Currency, currency.EUR)
		assert.Equal(t, category.Supercategory, 3)
	})
}
