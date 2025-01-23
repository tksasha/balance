package handlers_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
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

	createItemService := services.NewCreateItemService(itemRepository, categoryRepository)

	middleware := middlewares.NewCurrencyMiddleware().Wrap(
		handlers.NewCreateItemHandler(createItemService),
	)

	mux := http.NewServeMux()
	mux.Handle("POST /items", middleware)

	ctx := context.Background()

	t.Run("when form parsing error is happened, it should respond with 400", func(t *testing.T) {
		cleanup(ctx, t, db)

		body := strings.NewReader("%")

		request, err := http.NewRequestWithContext(ctx, http.MethodPost, "/items", body)
		assert.NilError(t, err)

		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("when input data is invalid it should render form", func(t *testing.T) {
		cleanup(ctx, t, db)

		request := newPostRequest(ctx, t, "/items", Params{"date": ""})

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})

	t.Run("when input data is valid, it should respond with 200", func(t *testing.T) {
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

		item := findItemByDate(usdContext(ctx, t), t, db.Connection, "2024-10-16")

		assert.Equal(t, item.GetDateAsString(), "2024-10-16")
		assert.Equal(t, item.CategoryID, 1101)
		assert.Equal(t, item.CategoryName, "Accoutrements")
		assert.Equal(t, item.Currency, currencies.USD)
		assert.Equal(t, item.Formula, "42.69+69.42")
		assert.Equal(t, item.GetSumAsString(), "112,11")
		assert.Equal(t, item.Description, "paper clips, notebooks, and pens")
	})
}
