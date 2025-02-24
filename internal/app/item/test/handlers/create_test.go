package handlers_test

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/tksasha/balance/internal/app/category"
	categoryrepository "github.com/tksasha/balance/internal/app/category/repository"
	categoryservice "github.com/tksasha/balance/internal/app/category/service"
	"github.com/tksasha/balance/internal/app/item/components"
	"github.com/tksasha/balance/internal/app/item/handlers"
	"github.com/tksasha/balance/internal/app/item/repository"
	"github.com/tksasha/balance/internal/app/item/service"
	"github.com/tksasha/balance/internal/common"
	commoncomponent "github.com/tksasha/balance/internal/common/component"
	"github.com/tksasha/balance/internal/common/currency"
	"github.com/tksasha/balance/internal/common/tests"
	"github.com/tksasha/balance/internal/db"
	nameprovider "github.com/tksasha/balance/internal/db/nameprovider/test"
	"gotest.tools/v3/assert"
)

func TestItemCreateHandler(t *testing.T) { //nolint:funlen
	handler, db := newCreateHandler(t)
	defer func() {
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	mux := mux(t, "POST /items", handler)

	ctx := t.Context()

	t.Run("responds 400 on parse form fails", func(t *testing.T) {
		request := tests.NewInvalidPostRequest(ctx, t, "/items")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("renders errors on invalid input", func(t *testing.T) {
		tests.Cleanup(ctx, t)

		params := tests.Params{"date": ""}

		request := tests.NewPostRequest(ctx, t, "/items", params)

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})

	t.Run("responds 204 when item created", func(t *testing.T) {
		tests.Cleanup(ctx, t)

		categoryToCreate := &category.Category{
			ID:       1101,
			Name:     "Accoutrements",
			Currency: currency.USD,
		}

		createCategory(t, db, categoryToCreate)

		params := tests.Params{
			"date":        "2024-10-16",
			"formula":     "42.69+69.42",
			"category_id": "1101",
			"description": "paper clips, notebooks, and pens",
		}

		request := tests.NewPostRequest(ctx, t, "/items?currency=usd", params)

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNoContent)

		item := tests.FindItemByDate(ctx, t, currency.USD, "2024-10-16")

		assert.Equal(t, item.Date.Format(time.DateOnly), "2024-10-16")
		assert.Equal(t, item.CategoryID, 1101)
		assert.Equal(t, item.CategoryName.String, "Accoutrements")
		assert.Equal(t, item.Currency, currency.USD)
		assert.Equal(t, item.Formula, "42.69+69.42")
		assert.Equal(t, item.Sum, 112.11)
		assert.Equal(t, item.Description, "paper clips, notebooks, and pens")
	})
}

func newCreateHandler(t *testing.T) (*handlers.CreateHandler, *sql.DB) {
	t.Helper()

	db := db.Open(t.Context(), nameprovider.New())

	itemRepository := repository.New(common.NewBaseRepository(), db)

	categoryRepository := categoryrepository.New(common.NewBaseRepository(), db)

	itemService := service.New(common.NewBaseService(), itemRepository, categoryRepository)

	categoryService := categoryservice.New(common.NewBaseService(), categoryRepository)

	itemComponent := components.NewItemsComponent(commoncomponent.New())

	handler := handlers.NewCreateHandler(common.NewBaseHandler(), itemService, categoryService, itemComponent)

	return handler, db
}
