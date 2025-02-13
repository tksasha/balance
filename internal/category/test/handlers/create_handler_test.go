package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/category/handlers"
	"github.com/tksasha/balance/internal/common/tests"
	"github.com/tksasha/balance/pkg/currencies"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCategoryCreateHandler(t *testing.T) {
	ctx := t.Context()

	service, db := tests.NewCategoryService(ctx, t)
	defer func() {
		_ = db.Close()
	}()

	mux := tests.NewMux(t, "POST /categories", handlers.NewCreateHandler(service))

	t.Run("responds 200 when input is invalid", func(t *testing.T) {
		params := tests.Params{"name": ""}

		request := tests.NewPostRequest(ctx, t, "/categories", params)

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		body := tests.GetResponseBody(t, recorder.Body)

		assert.Equal(t, recorder.Code, http.StatusOK)
		assert.Assert(t, is.Contains(body, "name: is required"))
	})

	t.Run("responds 201 on successful create", func(t *testing.T) {
		tests.Cleanup(ctx, t)

		params := tests.Params{
			"name":          "Miscellaneous",
			"income":        "true",
			"visible":       "true",
			"supercategory": "3",
		}

		request := tests.NewPostRequest(ctx, t, "/categories?currency=eur", params)

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusCreated)

		category := tests.FindCategoryByName(ctx, t, currencies.EUR, "Miscellaneous")

		assert.Equal(t, category.ID, 1)
		assert.Equal(t, category.Name, "Miscellaneous")
		assert.Equal(t, category.Income, true)
		assert.Equal(t, category.Visible, true)
		assert.Equal(t, category.Currency, currencies.EUR)
		assert.Equal(t, category.Supercategory, 3)
	})
}
