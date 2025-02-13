package handlers_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/tksasha/balance/internal/handlers"
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/pkg/currencies"
	"gotest.tools/v3/assert"
)

func TestCategoryListHandler(t *testing.T) {
	ctx := t.Context()

	service, db := newCategoryService(ctx, t)
	defer func() {
		_ = db.Close()
	}()

	mux := newMux(t, "GET /categories", handlers.NewCategoryListHandler(service))

	t.Run("responds 200 on no categories found", func(t *testing.T) {
		cleanup(ctx, t)

		request := newGetRequest(ctx, t, "/categories?currency=eur")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)
	})

	t.Run("responds 200 on categories found", func(t *testing.T) {
		cleanup(ctx, t)

		for id, name := range map[int]string{1: "category one", 2: "category two"} {
			categoryToCreate := &models.Category{
				ID:       id,
				Name:     name,
				Currency: currencies.EUR,
				Visible:  true,
			}

			createCategory(ctx, t, categoryToCreate)
		}

		request := newGetRequest(ctx, t, "/categories?currency=eur")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		response, err := io.ReadAll(recorder.Body)
		if err != nil {
			t.Fatalf("failed to read response body, error: %v", err)
		}

		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Assert(t, strings.Contains(string(response), "category one"))
		assert.Assert(t, strings.Contains(string(response), "category two"))
	})
}
