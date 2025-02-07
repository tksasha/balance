package handlers_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/apperrors"
	"github.com/tksasha/balance/internal/handlers"
	"github.com/tksasha/balance/internal/middlewares"
	mocksforhandlers "github.com/tksasha/balance/mocks/handlers"
	"github.com/tksasha/balance/pkg/currencies"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestItemEditHandler(t *testing.T) {
	controller := gomock.NewController(t)

	itemService := mocksforhandlers.NewMockItemService(controller)

	middleware := middlewares.NewCurrencyMiddleware().Wrap(
		handlers.NewItemEditHandler(itemService),
	)

	ctx := context.Background()

	ctxWithValue := context.WithValue(ctx, currencies.CurrencyContextValue{}, currencies.USD)

	mux := http.NewServeMux()
	mux.Handle("GET /items/{id}/edit", middleware)

	t.Run("responds 404 on no item found", func(t *testing.T) {
		itemService.
			EXPECT().
			GetItem(ctxWithValue, "1514").
			Return(nil, apperrors.ErrResourceNotFound)

		request := newGetRequest(ctx, t, "/items/1514/edit?currency=usd")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusNotFound, recorder.Code)
	})
}
