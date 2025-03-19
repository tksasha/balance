package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	categoryrepository "github.com/tksasha/balance/internal/app/category/repository"
	"github.com/tksasha/balance/internal/app/item/handlers"
	"github.com/tksasha/balance/internal/app/item/repository"
	"github.com/tksasha/balance/internal/app/item/service"
	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/db/nameprovider"
	"gotest.tools/v3/assert"
	"gotest.tools/v3/golden"
)

func TestItemIndexHandler(t *testing.T) {
	ctx := t.Context()

	db := db.Open(ctx, nameprovider.NewTestProvider())
	defer func() {
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	itemRepository := repository.New(db)

	categoryRepository := categoryrepository.New(db)

	itemService := service.New(itemRepository, categoryRepository)

	handler := handlers.NewIndexHandler(itemService)

	mux := mux(t, "GET /items", handler)

	t.Run("responds 200 on items found", func(t *testing.T) {
		request, err := http.NewRequestWithContext(ctx, http.MethodGet, "/items?currency=eur&month=12&year=2025", nil)
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)

		golden.Assert(t, recorder.Header().Get("Hx-Trigger-After-Swap"),
			"index-hx-trigger-after-swap-header.json")
	})
}
