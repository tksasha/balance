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

func TestCreateItemHandler_ServeHTTP(t *testing.T) { //nolint:funlen
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

	t.Run("responds with 400 when parse form fails", func(t *testing.T) {
		cleanup(ctx, t, db)

		request := newInvalidPostRequest(ctx, t, "/items")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("renders form when input data is invalid", func(t *testing.T) {
		cleanup(ctx, t, db)

		request := newPostRequest(ctx, t, "/items", Params{"date": ""})

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})

	t.Run("responds with 200 when there no errors", func(t *testing.T) {
		cleanup(ctx, t, db)

		createCategory(ctx, t, db,
			&models.Category{
				ID:       1101,
				Name:     "Accoutrements",
				Currency: currencies.USD,
			},
		)

		params := Params{
			"date":        "2024-10-16",
			"formula":     "42.69+69.42",
			"category_id": "1101",
			"description": "paper clips, notebooks, and pens",
		}

		request := newPostRequest(ctx, t, "/items?currency=usd", params)

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)

		item := findItemByDate(usdContext(ctx, t), t, db, "2024-10-16")

		assert.Equal(t, item.GetDateAsString(), "2024-10-16")
		assert.Equal(t, item.CategoryID, 1101)
		assert.Equal(t, item.CategoryName, "Accoutrements")
		assert.Equal(t, item.Currency, currencies.USD)
		assert.Equal(t, item.Formula, "42.69+69.42")
		assert.Equal(t, item.GetSumAsString(), "112,11")
		assert.Equal(t, item.Description, "paper clips, notebooks, and pens")
	})
}
