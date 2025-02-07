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

func TestItemUpdateHandler(t *testing.T) { //nolint:funlen
	dbNameProvider := providers.NewDBNameProvider()

	db := db.Open(dbNameProvider)

	itemRepository := repositories.NewItemRepository(db)
	categoryRepository := repositories.NewCategoryRepository(db)

	itemService := services.NewItemService(itemRepository, categoryRepository)

	middleware := middlewares.NewCurrencyMiddleware().Wrap(
		handlers.NewItemUpdateHandler(itemService),
	)

	ctx := context.Background()

	mux := http.NewServeMux()
	mux.Handle("PATCH /items/{id}", middleware)

	t.Run("responds 400 on invalid input", func(t *testing.T) {
		cleanup(ctx, t, db)

		request := newInvalidPatchRequest(ctx, t, "/items/1138")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusBadRequest)
	})

	t.Run("responds 404 on no item found", func(t *testing.T) {
		cleanup(ctx, t, db)

		request := newPatchRequest(ctx, t, "/items/1218", Params{"date": "2025-01-25"})

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNotFound)
	})

	t.Run("responds 200 on successful update", func(t *testing.T) {
		cleanup(ctx, t, db)

		createCategory(ctx, t, db,
			&models.Category{
				ID:       1148,
				Name:     "Pharmaceutical",
				Currency: currencies.EUR,
			},
		)

		createItem(ctx, t, db,
			&models.Item{
				ID:         1143,
				CategoryID: 1148,
				Currency:   currencies.EUR,
			},
		)

		request := newPatchRequest(ctx, t, "/items/1143?currency=eur",
			Params{
				"date":        "2025-01-25",
				"formula":     "24 + 11 + 49",
				"category_id": "1148",
				"description": "pizza, ninja and disco",
			},
		)

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)

		item := findItemByDate(eurContext(ctx, t), t, db, "2025-01-25")

		assert.Equal(t, item.ID, 1143)
	})
}
