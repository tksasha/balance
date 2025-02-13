package handlers_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/handlers"
	"github.com/tksasha/balance/pkg/currencies"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCategoryCreateHandler(t *testing.T) {
	ctx := t.Context()

	service, db := newCategoryService(ctx, t)
	defer func() {
		_ = db.Close()
	}()

	mux := newMux(t, "POST /categories", handlers.NewCategoryCreateHandler(service))

	t.Run("responds 200 on invalid input", func(t *testing.T) {
		request := newPostRequest(ctx, t, "/categories", Params{"name": ""})

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		response, err := io.ReadAll(recorder.Body)
		if err != nil {
			t.Fatalf("failed to read response body, error: %v", err)
		}

		assert.Equal(t, recorder.Code, http.StatusOK)
		assert.Assert(t, is.Contains(string(response), "name: is required"))
	})

	t.Run("responds 200 on successful create", func(t *testing.T) {
		cleanup(ctx, t)

		request := newPostRequest(ctx, t, "/categories?currency=eur",
			Params{
				"name":          "Miscellaneous",
				"income":        "true",
				"visible":       "true",
				"supercategory": "3",
			},
		)

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		response, err := io.ReadAll(recorder.Body)
		if err != nil {
			t.Fatalf("failed to read response body, error: %v", err)
		}

		assert.Equal(t, recorder.Code, http.StatusOK)
		assert.Assert(t, is.Contains(string(response), "create category page"))

		category := findCategoryByName(ctx, t, currencies.EUR, "Miscellaneous")

		assert.Equal(t, category.ID, 1)
		assert.Equal(t, category.Name, "Miscellaneous")
		assert.Equal(t, category.Income, true)
		assert.Equal(t, category.Visible, true)
		assert.Equal(t, category.Currency, currencies.EUR)
		assert.Equal(t, category.Supercategory, 3)
	})
}
