package handlers_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/handlers"
	"github.com/tksasha/balance/internal/middlewares"
	"github.com/tksasha/balance/internal/models"
	mocksforhandlers "github.com/tksasha/balance/mocks/handlers"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestGetItemHandler(t *testing.T) {
	controller := gomock.NewController(t)

	itemService := mocksforhandlers.NewMockItemService(controller)

	middleware := middlewares.NewCurrencyMiddleware().Wrap(
		handlers.NewGetItemHandler(itemService),
	)

	ctx := context.Background()

	ctxWithValue := context.WithValue(ctx, models.CurrencyContextValue{}, models.USD)

	t.Run("when get item by id returns an error, it should respond with 500 code", func(t *testing.T) {
		itemService.
			EXPECT().
			GetItem(ctxWithValue, "1514").
			Return(nil, errors.New("get item by id error"))

		request, err := http.NewRequestWithContext(ctx, http.MethodGet, "/items/1514?currency=usd", nil)
		assert.NilError(t, err)

		route := http.NewServeMux()
		route.Handle("GET /items/{id}", middleware)

		response := httptest.NewRecorder()

		route.ServeHTTP(response, request)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})
}
