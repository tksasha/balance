package handlers_test

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/tksasha/balance/internal/handlers"
	"github.com/tksasha/balance/internal/middlewares"
	"github.com/tksasha/balance/internal/models"
	mocksforhandlers "github.com/tksasha/balance/mocks/handlers"
	"github.com/tksasha/balance/pkg/currencies"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestGetItemsHandler(t *testing.T) {
	controller := gomock.NewController(t)

	itemService := mocksforhandlers.NewMockItemService(controller)

	middleware := middlewares.NewCurrencyMiddleware().Wrap(
		handlers.NewGetItemsHandler(itemService),
	)

	ctx := context.Background()

	ctxWithValue := context.WithValue(ctx, currencies.CurrencyContextValue{}, currencies.EUR)

	t.Run("when GetItems returns an error, it should respond with 500 code", func(t *testing.T) {
		itemService.
			EXPECT().
			GetItems(ctxWithValue).
			Return(nil, errors.New("get items error"))

		request, err := http.NewRequestWithContext(ctx, http.MethodGet, "/items?currency=eur", nil)
		assert.NilError(t, err)

		recorder := httptest.NewRecorder()

		middleware.ServeHTTP(recorder, request)

		body, err := io.ReadAll(recorder.Body)
		assert.NilError(t, err)

		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
		assert.Assert(t, strings.Contains(string(body), "Internal Server Error"))
	})

	t.Run("when GetItems doesn't return any error, it should respond with 200 code", func(t *testing.T) {
		itemService.
			EXPECT().
			GetItems(ctxWithValue).
			Return(models.Items{}, nil)

		request, err := http.NewRequestWithContext(ctx, http.MethodGet, "/items?currency=eur", nil)
		assert.NilError(t, err)

		recorder := httptest.NewRecorder()

		middleware.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}
