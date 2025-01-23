package handlers_test

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/handlers"
	"github.com/tksasha/balance/internal/middlewares"
	providers "github.com/tksasha/balance/internal/providers/test"
	"github.com/tksasha/balance/internal/repositories"
	"github.com/tksasha/balance/internal/services"
	"github.com/tksasha/balance/pkg/currencies"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCreateCategoryHandlerTest(t *testing.T) { //nolint:funlen
	dbNameProvider := providers.NewDBNameProvider()

	db := db.Open(dbNameProvider)

	categoryRepository := repositories.NewCategoryRepository(db)

	categoryService := services.NewCategoryService(categoryRepository)

	middleware := middlewares.NewCurrencyMiddleware().Wrap(
		handlers.NewCreateCategoryHandler(categoryService),
	)

	mux := http.NewServeMux()
	mux.Handle("POST /categories", middleware)

	ctx := context.Background()

	t.Run("when input data is invalid, it should render validation errors", func(t *testing.T) {
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

	t.Run("when input data is valid, it should create category", func(t *testing.T) {
		cleanup(ctx, t, db)

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

		category := findCategoryByName(eurContext(ctx, t), t, db, "Miscellaneous")

		assert.Equal(t, category.ID, 1)
		assert.Equal(t, category.Name, "Miscellaneous")
		assert.Equal(t, category.Income, true)
		assert.Equal(t, category.Visible, true)
		assert.Equal(t, category.Currency, currencies.EUR)
		assert.Equal(t, category.Supercategory, 3)
	})
}
