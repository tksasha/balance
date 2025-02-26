package handlers_test

import (
	"database/sql"
	"io"
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
	"github.com/tksasha/balance/internal/common/currency"
	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/db/nameprovider"
	"gotest.tools/v3/assert"
	"gotest.tools/v3/golden"
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
		request, err := http.NewRequestWithContext(ctx, http.MethodGet, "/items/1514/edit?currency=usd", nil)
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusNotFound, recorder.Code)
	})

	t.Run("responds 200 when item found", func(t *testing.T) {
		cleanup(t, db)

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

		createItem(t, db, itemToCreate)

		request, err := http.NewRequestWithContext(ctx, http.MethodGet, "/items/1745/edit", nil)
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)

		response, err := io.ReadAll(recorder.Body)
		if err != nil {
			t.Fatal(err)
		}

		golden.Assert(t, string(response), "edit.html")
	})
}

func newEditHandler(t *testing.T) (*handlers.EditHandler, *sql.DB) {
	t.Helper()

	db := db.Open(t.Context(), nameprovider.NewTestProvider())

	itemRepository := repository.New(db)

	categoryRepository := categoryrepository.New(db)

	itemService := service.New(itemRepository, categoryRepository)

	categoryService := categoryservice.New(categoryRepository)

	itemComponent := components.NewItemsComponent()

	handler := handlers.NewEditHandler(itemService, categoryService, itemComponent)

	return handler, db
}
