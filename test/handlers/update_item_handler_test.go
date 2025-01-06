package handlers_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/handlers"
	"github.com/tksasha/balance/internal/middlewares"
	providers "github.com/tksasha/balance/internal/providers/test"
	"github.com/tksasha/balance/internal/repositories"
	"github.com/tksasha/balance/internal/services"
	"gotest.tools/v3/assert"
)

func TestUpdateItemHandler(t *testing.T) {
	dbNameProvider := providers.NewDBNameProvider()

	dbConnection := db.Open(dbNameProvider)

	itemRepository := repositories.NewItemRepository(dbConnection)

	itemService := services.NewItemService(itemRepository)

	middleware := middlewares.NewCurrencyMiddleware().Wrap(
		handlers.NewUpdateItemHandler(itemService),
	)

	ctx := context.Background()

	route := http.NewServeMux()
	route.Handle("PATCH /items/{id}", middleware)

	t.Run("when item is not found, it should respond with 404", func(t *testing.T) {
		request, err := http.NewRequestWithContext(ctx, http.MethodPatch, "/items/1218", nil)
		assert.NilError(t, err)

		recorder := httptest.NewRecorder()

		route.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusNotFound, recorder.Code)
	})
}
