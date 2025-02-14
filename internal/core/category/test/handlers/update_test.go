package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/core/category/handlers"
	"github.com/tksasha/balance/internal/core/common/tests"
	"github.com/tksasha/balance/pkg/currencies"
	"gotest.tools/v3/assert"
)

func TestCategoryUpdateHandler(t *testing.T) { //nolint:funlen
	ctx := t.Context()

	service, db := tests.NewCategoryService(ctx, t)
	defer func() {
		_ = db.Close()
	}()

	mux := tests.NewMux(t, "PATCH /categories/{id}", handlers.NewUpdateHandler(service))

	t.Run("responds 404 on no category found", func(t *testing.T) {
		tests.Cleanup(ctx, t)

		request := tests.NewPatchRequest(ctx, t, "/categories/1141", nil)

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNotFound)
	})

	t.Run("renders errors when validation failed", func(t *testing.T) {
		tests.Cleanup(ctx, t)

		for id, name := range map[int]string{1151: "Heterogeneous", 11654: "Paraphernalia"} {
			categoryToCreate := &category.Category{
				ID:       id,
				Name:     name,
				Currency: currencies.USD,
			}

			tests.CreateCategory(ctx, t, categoryToCreate)
		}

		params := tests.Params{"name": "Paraphernalia"}

		request := tests.NewPatchRequest(ctx, t, "/categories/1151?currency=usd", params)

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		body := tests.GetResponseBody(t, recorder.Body)

		assert.Equal(t, recorder.Code, http.StatusOK)
		assert.Assert(t, strings.Contains(body, "name: already exists"))
	})

	t.Run("responds 204 when category updated", func(t *testing.T) {
		tests.Cleanup(ctx, t)

		categoryToCreate := &category.Category{
			ID:            1208,
			Name:          "Paraphernalia",
			Income:        false,
			Visible:       false,
			Currency:      currencies.USD,
			Supercategory: 5,
		}

		tests.CreateCategory(ctx, t, categoryToCreate)

		params := tests.Params{
			"name":          "Heterogeneous",
			"income":        "true",
			"visible":       "true",
			"supercategory": "4",
		}

		request := tests.NewPatchRequest(ctx, t, "/categories/1208?currency=usd", params)

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNoContent)

		category := tests.FindCategoryByID(ctx, t, currencies.USD, 1208)

		assert.Equal(t, category.ID, 1208)
		assert.Equal(t, category.Name, "Heterogeneous")
		assert.Equal(t, category.Income, true)
		assert.Equal(t, category.Visible, true)
		assert.Equal(t, category.Currency, currencies.USD)
		assert.Equal(t, category.Supercategory, 4)
	})
}
