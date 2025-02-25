package handlers_test

import (
	"database/sql"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/tksasha/balance/internal/backoffice/category"
	"github.com/tksasha/balance/internal/backoffice/category/component"
	"github.com/tksasha/balance/internal/backoffice/category/handlers"
	"github.com/tksasha/balance/internal/backoffice/category/repository"
	"github.com/tksasha/balance/internal/backoffice/category/service"
	commonComponent "github.com/tksasha/balance/internal/common/component"
	"github.com/tksasha/balance/internal/common/currency"
	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/db/nameprovider"
	"gotest.tools/v3/assert"
)

func TestCategoryUpdateHandler(t *testing.T) { //nolint:funlen
	handler, db := newUpdateHandler(t)
	defer func() {
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	mux := mux(t, "PATCH /backoffice/categories/{id}", handler)

	ctx := t.Context()

	t.Run("responds with 400 when request is invalid", func(t *testing.T) {
		request, err := http.NewRequestWithContext(ctx, http.MethodPatch, "/backoffice/categories/1635", nil)
		if err != nil {
			t.Fatal(err)
		}

		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusBadRequest)
	})

	t.Run("responds 404 on no category found", func(t *testing.T) {
		formData := url.Values{"name": {""}}

		body := strings.NewReader(formData.Encode())

		request, err := http.NewRequestWithContext(ctx, http.MethodPatch, "/backoffice/categories/1141", body)
		if err != nil {
			t.Fatal(err)
		}

		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNotFound)
	})

	t.Run("renders errors when validation failed", func(t *testing.T) {
		cleanup(t, db)

		for id, name := range map[int]string{1151: "Heterogeneous", 11654: "Paraphernalia"} {
			categoryToCreate := &category.Category{
				ID:       id,
				Name:     name,
				Currency: currency.USD,
			}

			createCategory(t, db, categoryToCreate)
		}

		values := url.Values{"name": {"Paraphernalia"}}

		body := strings.NewReader(values.Encode())

		request, err := http.NewRequestWithContext(
			ctx,
			http.MethodPatch,
			"/backoffice/categories/1151?currency=usd",
			body,
		)
		if err != nil {
			t.Fatal(err)
		}

		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)

		responseBody, err := io.ReadAll(recorder.Body)
		if err != nil {
			t.Fatal(err)
		}

		assert.Assert(t, strings.Contains(string(responseBody), "name: already exists"))
	})

	t.Run("responds 204 when category updated", func(t *testing.T) {
		cleanup(t, db)

		categoryToCreate := &category.Category{
			ID:            1208,
			Name:          "Paraphernalia",
			Income:        false,
			Visible:       false,
			Currency:      currency.USD,
			Supercategory: 5,
		}

		createCategory(t, db, categoryToCreate)

		formData := url.Values{
			"name":          {"Heterogeneous"},
			"income":        {"true"},
			"visible":       {"true"},
			"supercategory": {"4"},
		}

		body := strings.NewReader(formData.Encode())

		request, err := http.NewRequestWithContext(
			ctx,
			http.MethodPatch,
			"/backoffice/categories/1208?currency=usd",
			body,
		)
		if err != nil {
			t.Fatal(err)
		}

		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNoContent)

		category := findCategoryByID(t, db, currency.USD, 1208)

		assert.Equal(t, category.ID, 1208)
		assert.Equal(t, category.Name, "Heterogeneous")
		assert.Equal(t, category.Income, true)
		assert.Equal(t, category.Visible, true)
		assert.Equal(t, category.Currency, currency.USD)
		assert.Equal(t, category.Supercategory, 4)
	})
}

func newUpdateHandler(t *testing.T) (*handlers.UpdateHandler, *sql.DB) {
	t.Helper()

	db := db.Open(t.Context(), nameprovider.NewTestProvider())

	categoryRepository := repository.New(db)

	categoryService := service.New(categoryRepository)

	categoryComponent := component.New(commonComponent.New())

	handler := handlers.NewUpdateHandler(categoryService, categoryComponent)

	return handler, db
}
