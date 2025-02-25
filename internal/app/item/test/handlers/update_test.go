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

func TestItemUpdateHandler(t *testing.T) { //nolint:funlen
	handler, db := newUpdateHandler(t)
	defer func() {
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	mux := tests.NewMux(t, "PATCH /items/{id}", handler)

	ctx := t.Context()

	t.Run("responds 400 on invalid input", func(t *testing.T) {
		tests.Cleanup(ctx, t)

		request := tests.NewInvalidPatchRequest(ctx, t, "/items/1138")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusBadRequest)
	})

	t.Run("responds 404 on no item found", func(t *testing.T) {
		tests.Cleanup(ctx, t)

		params := tests.Params{"date": "2025-01-25"}

		request := tests.NewPatchRequest(ctx, t, "/items/1218", params)

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNotFound)
	})

	t.Run("responds 204 on successful update", func(t *testing.T) {
		tests.Cleanup(ctx, t)

		categoryToCreate := &category.Category{
			ID:       1148,
			Name:     "Pharmaceutical",
			Currency: currency.EUR,
		}

		createCategory(t, db, categoryToCreate)

		itemToCreate := &item.Item{
			ID:         1143,
			CategoryID: 1148,
			Currency:   currency.EUR,
		}

		tests.CreateItem(ctx, t, itemToCreate)

		params := tests.Params{
			"date":        "2025-01-25",
			"formula":     "24 + 11 + 49",
			"category_id": "1148",
			"description": "pizza, ninja and disco",
		}

		request := tests.NewPatchRequest(ctx, t, "/items/1143?currency=eur", params)

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNoContent)

		item := tests.FindItemByDate(ctx, t, currency.EUR, "2025-01-25")

		assert.Equal(t, item.ID, 1143)
	})
}

func newUpdateHandler(t *testing.T) (*handlers.UpdateHandler, *sql.DB) {
	t.Helper()

	db := db.Open(t.Context(), nameprovider.New())

	itemRepository := repository.New(db)

	categoryRepository := categoryrepository.New(db)

	itemService := service.New(common.NewBaseService(), itemRepository, categoryRepository)

	categoryService := categoryservice.New(common.NewBaseService(), categoryRepository)

	itemComponent := components.NewItemsComponent(commoncomponent.New())

	handler := handlers.NewUpdateHandler(itemService, categoryService, itemComponent)

	return handler, db
}
