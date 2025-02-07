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
	"github.com/tksasha/balance/test/testutils"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestItemListHandler(t *testing.T) {
	controller := gomock.NewController(t)

	itemService := mocksforhandlers.NewMockItemService(controller)

	middleware := middlewares.NewCurrencyMiddleware().Wrap(
		handlers.NewItemListHandler(itemService),
	)

	ctx := context.Background()

	ctxWithValue := context.WithValue(ctx, currencies.CurrencyContextValue{}, currencies.EUR)

	mux := http.NewServeMux()
	mux.Handle("GET /items", middleware)

	t.Run("responds 500 on internal server error", func(t *testing.T) {
		itemService.
			EXPECT().
			GetItems(ctxWithValue).
			Return(nil, errors.New("get items error"))

		request := testutils.NewGetRequest(ctx, t, "/items?currency=eur")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		body, err := io.ReadAll(recorder.Body)
		assert.NilError(t, err)

		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
		assert.Assert(t, strings.Contains(string(body), "Internal Server Error"))
	})

	t.Run("responds 200 on items found", func(t *testing.T) {
		itemService.
			EXPECT().
			GetItems(ctxWithValue).
			Return(models.Items{}, nil)

		request := testutils.NewGetRequest(ctx, t, "/items?currency=eur")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}
