package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/app/category"
	categoryrepository "github.com/tksasha/balance/internal/app/category/repository"
	"github.com/tksasha/balance/internal/app/item"
	"github.com/tksasha/balance/internal/app/item/handlers"
	"github.com/tksasha/balance/internal/app/item/repository"
	"github.com/tksasha/balance/internal/app/item/service"
	"github.com/tksasha/balance/internal/common/currency"
	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/db/nameprovider"
	"github.com/tksasha/balance/internal/server/middlewares"
	"gotest.tools/v3/assert"
	"gotest.tools/v3/golden"
)

func TestDeleteItemHandler(t *testing.T) { //nolint:funlen
	db := db.Open(t.Context(), nameprovider.NewTestProvider())
	defer func() {
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	itemRepository := repository.New(db)

	categoryRepository := categoryrepository.New(db)

	itemService := service.New(itemRepository, categoryRepository)

	handler := handlers.NewDeleteHandler(itemService)

	mux := http.NewServeMux()

	next := http.Handler(handler)

	for _, middleware := range middlewares.New() {
		next = middleware.Wrap(next)
	}

	mux.Handle("DELETE /items/{id}", next)

	ctx := t.Context()

	t.Run("renders 404 when item is not found", func(t *testing.T) {
		request, err := http.NewRequestWithContext(ctx, http.MethodDelete, "/items/1043", nil)
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNotFound)
	})

	t.Run("render 200 when item is deleted", func(t *testing.T) {
		cleanup(t, db)

		categoryToCreate := &category.Category{
			ID:       1047,
			Currency: currency.UAH,
		}

		createCategory(t, db, categoryToCreate)

		itemToCreate := &item.Item{
			ID:         1045,
			Currency:   currency.UAH,
			CategoryID: 1047,
			Date:       date(t, "2025-03-10"),
		}

		createItem(t, db, itemToCreate)

		request, err := http.NewRequestWithContext(ctx, http.MethodDelete, "/items/1045", nil)
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)

		golden.Assert(t, recorder.Header().Get("Hx-Trigger-After-Swap"),
			"delete-hx-trigger-after-swap-header.json")
	})
}
