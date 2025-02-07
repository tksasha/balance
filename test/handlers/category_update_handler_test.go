package handlers_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/handlers"
	"github.com/tksasha/balance/internal/middlewares"
	"github.com/tksasha/balance/internal/models"
	providers "github.com/tksasha/balance/internal/providers/test"
	"github.com/tksasha/balance/internal/repositories"
	"github.com/tksasha/balance/internal/services"
	"github.com/tksasha/balance/pkg/currencies"
	"gotest.tools/v3/assert"
)

func TestCategoryUpdateHandler(t *testing.T) { //nolint:funlen
	dbNameProvider := providers.NewDBNameProvider()

	db := db.Open(dbNameProvider)

	categoryRepository := repositories.NewCategoryRepository(db)

	categoryService := services.NewCategoryService(categoryRepository)

	middleware := middlewares.NewCurrencyMiddleware().Wrap(
		handlers.NewCategoryUpdateHandler(categoryService),
	)

	mux := http.NewServeMux()
	mux.Handle("PATCH /categories/{id}", middleware)

	ctx := context.Background()

	t.Run("responds 404 on no category found", func(t *testing.T) {
		cleanup(ctx, t, db)

		request := newPatchRequest(ctx, t, "/categories/1141", nil)

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNotFound)
	})

	t.Run("responds 200 on duplicates name", func(t *testing.T) {
		cleanup(ctx, t, db)

		for id, name := range map[int]string{1151: "Heterogeneous", 11654: "Paraphernalia"} {
			createCategory(ctx, t, db,
				&models.Category{
					ID:       id,
					Name:     name,
					Currency: currencies.USD,
				},
			)
		}

		request := newPatchRequest(ctx, t,
			"/categories/1151?currency=usd",
			Params{"name": "Paraphernalia"},
		)

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)
	})

	t.Run("responds 200 on successful update", func(t *testing.T) {
		cleanup(ctx, t, db)

		createCategory(ctx, t, db,
			&models.Category{
				ID:            1208,
				Name:          "Paraphernalia",
				Income:        false,
				Visible:       false,
				Currency:      currencies.USD,
				Supercategory: 5,
			},
		)

		request := newPatchRequest(ctx, t, "/categories/1208?currency=usd",
			Params{
				"name":          "Heterogeneous",
				"income":        "true",
				"visible":       "true",
				"supercategory": "4",
			},
		)

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)

		category := findCategoryByID(usdContext(ctx, t), t, db, 1208)

		assert.Equal(t, category.ID, 1208)
		assert.Equal(t, category.Name, "Heterogeneous")
		assert.Equal(t, category.Income, true)
		assert.Equal(t, category.Visible, true)
		assert.Equal(t, category.Currency, currencies.USD)
		assert.Equal(t, category.Supercategory, 4)
	})
}
