package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/tksasha/balance/internal/category"
	"github.com/tksasha/balance/internal/category/handlers"
	"github.com/tksasha/balance/internal/common/testutils"
	"github.com/tksasha/balance/pkg/currencies"
	"gotest.tools/v3/assert"
)

func TestCategoryListHandler(t *testing.T) {
	ctx := t.Context()

	service, db := testutils.NewCategoryService(ctx, t)
	defer func() {
		_ = db.Close()
	}()

	mux := testutils.NewMux(t, "GET /categories", handlers.NewListHandler(service))

	t.Run("responds 200 on no categories found", func(t *testing.T) {
		testutils.Cleanup(ctx, t)

		request := testutils.NewGetRequest(ctx, t, "/categories?currency=eur")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)
	})

	t.Run("responds 200 on categories found", func(t *testing.T) {
		testutils.Cleanup(ctx, t)

		for id, name := range map[int]string{1: "category one", 2: "category two"} {
			categoryToCreate := &category.Category{
				ID:       id,
				Name:     name,
				Currency: currencies.EUR,
				Visible:  true,
			}

			testutils.CreateCategory(ctx, t, categoryToCreate)
		}

		request := testutils.NewGetRequest(ctx, t, "/categories?currency=eur")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		body := testutils.GetResponseBody(t, recorder.Body)

		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Assert(t, strings.Contains(body, "category one"))
		assert.Assert(t, strings.Contains(body, "category two"))
	})
}
