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
	mockedservices "github.com/tksasha/balance/mocks/services"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestGetCategoriesHandler_GetCategories(t *testing.T) {
	controller := gomock.NewController(t)

	categoryService := mockedservices.NewMockCategoryService(controller)

	middleware := middlewares.NewCurrencyMiddleware().Wrap(
		handlers.NewGetCategoriesHandler(categoryService),
	)

	ctx := context.Background()

	ctxWithValue := context.WithValue(ctx, models.CurrencyContextValue{}, models.EUR)

	t.Run("when GetCategories returns an error it should respond with 500 code", func(t *testing.T) {
		categoryService.EXPECT().GetCategories(ctxWithValue).Return(nil, errors.New("get categories error"))

		request, err := http.NewRequestWithContext(ctx, http.MethodGet, "/categories?currency=eur", nil)
		assert.NilError(t, err)

		recorder := httptest.NewRecorder()

		middleware.ServeHTTP(recorder, request)

		body, err := io.ReadAll(recorder.Body)
		assert.NilError(t, err)

		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
		assert.Assert(t, strings.Contains(string(body), "Internal Server Error"))
	})

	t.Run("when GetCategories doesn't return any error it should respond with 200 code", func(t *testing.T) {
		categoryService.EXPECT().GetCategories(ctxWithValue).Return(models.Categories{}, nil)

		request, err := http.NewRequestWithContext(ctx, http.MethodGet, "/categories?currency=eur", nil)
		assert.NilError(t, err)

		recorder := httptest.NewRecorder()

		middleware.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}
