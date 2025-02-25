package handlers_test

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/app/category"
	categoryrepository "github.com/tksasha/balance/internal/app/category/repository"
	categoryservice "github.com/tksasha/balance/internal/app/category/service"
	"github.com/tksasha/balance/internal/app/item"
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

func TestItemEditHandler(t *testing.T) {
	handler, db := newEditHandler(t)
	defer func() {
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	mux := mux(t, "GET /items/{id}/edit", handler)

	ctx := t.Context()

	t.Run("responds 404 when item not found", func(t *testing.T) {
		request := tests.NewGetRequest(ctx, t, "/items/1514/edit?currency=usd")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusNotFound, recorder.Code)
	})

	t.Run("responds 200 when item found", func(t *testing.T) {
		tests.Cleanup(ctx, t)

		categoryToCreate := &category.Category{
			ID:       5,
			Currency: currency.UAH,
		}

		createCategory(t, db, categoryToCreate)

		itemToCreate := &item.Item{
			ID:         1745,
			Currency:   currency.UAH,
			CategoryID: 5,
		}

		tests.CreateItem(ctx, t, itemToCreate)

		request := tests.NewGetRequest(ctx, t, "/items/1745/edit")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)
	})
}

func newEditHandler(t *testing.T) (*handlers.EditHandler, *sql.DB) {
	t.Helper()

	db := db.Open(t.Context(), nameprovider.New())

	itemRepository := repository.New(common.NewBaseRepository(), db)

	categoryRepository := categoryrepository.New(common.NewBaseRepository(), db)

	itemService := service.New(common.NewBaseService(), itemRepository, categoryRepository)

	categoryService := categoryservice.New(common.NewBaseService(), categoryRepository)

	itemComponent := components.NewItemsComponent(commoncomponent.New())

	handler := handlers.NewEditHandler(itemService, categoryService, itemComponent)

	return handler, db
}
