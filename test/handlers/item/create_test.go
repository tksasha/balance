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
	"github.com/tksasha/balance/test/testutils"
	"gotest.tools/v3/assert"
)

func TestCreate(t *testing.T) { //nolint:funlen
	dbNameProvider := providers.NewDBNameProvider()

	db := db.Open(dbNameProvider)

	itemRepository := repositories.NewItemRepository(db)
	categoryRepository := repositories.NewCategoryRepository(db)

	itemService := services.NewItemService(itemRepository, categoryRepository)

	middleware := middlewares.NewCurrencyMiddleware().Wrap(
		handlers.NewCreateItemHandler(itemService),
	)

	mux := http.NewServeMux()
	mux.Handle("POST /items", middleware)

	ctx := context.Background()

	t.Run("responds 400 on parse form fails", func(t *testing.T) {
		request := testutils.NewInvalidPostRequest(ctx, t, "/items")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("renders errors on invalid input", func(t *testing.T) {
		testutils.Cleanup(ctx, t, db)

		request := testutils.NewPostRequest(ctx, t, "/items", testutils.Params{"date": ""})

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})

	t.Run("responds 200 on succcessful create", func(t *testing.T) {
		testutils.Cleanup(ctx, t, db)

		testutils.CreateCategory(ctx, t, db,
			&models.Category{
				ID:       1101,
				Name:     "Accoutrements",
				Currency: currencies.USD,
			},
		)

		params := testutils.Params{
			"date":        "2024-10-16",
			"formula":     "42.69+69.42",
			"category_id": "1101",
			"description": "paper clips, notebooks, and pens",
		}

		request := testutils.NewPostRequest(ctx, t, "/items?currency=usd", params)

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)

		item := testutils.FindItemByDate(testutils.USDContext(ctx, t), t, db, "2024-10-16")

		assert.Equal(t, item.GetDateAsString(), "2024-10-16")
		assert.Equal(t, item.CategoryID, 1101)
		assert.Equal(t, item.CategoryName, "Accoutrements")
		assert.Equal(t, item.Currency, currencies.USD)
		assert.Equal(t, item.Formula, "42.69+69.42")
		assert.Equal(t, item.GetSumAsString(), "112,11")
		assert.Equal(t, item.Description, "paper clips, notebooks, and pens")
	})
}
