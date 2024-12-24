package handlers_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/tksasha/balance/internal/handlers"
	"github.com/tksasha/balance/internal/models"
	mocksforhandlers "github.com/tksasha/balance/mocks/handlers"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestCreateItemHandler_ServeHTTP(t *testing.T) { //nolint:funlen
	controller := gomock.NewController(t)

	itemService := mocksforhandlers.NewMockItemService(controller)
	categoryService := mocksforhandlers.NewMockCategoryService(controller)

	handler := handlers.NewCreateItemHandler(itemService, categoryService)

	ctx := context.Background()

	t.Run("when form parsing error is happened it should respond with 400", func(t *testing.T) {
		body := strings.NewReader("%")

		request, err := http.NewRequestWithContext(ctx, http.MethodPost, "/items", body)
		assert.NilError(t, err)

		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		recorder := httptest.NewRecorder()

		handler.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("when input data is invalid it should render form", func(t *testing.T) {
		formData := url.Values{}
		formData.Set("date", "")

		body := strings.NewReader(formData.Encode())

		request, err := http.NewRequestWithContext(ctx, http.MethodPost, "/items", body)
		assert.NilError(t, err)

		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		categoryService.
			EXPECT().
			GetCategories(ctx).
			Return(models.Categories{}, nil)

		recorder := httptest.NewRecorder()

		handler.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})

	t.Run("when item creator returns an error it should respond with 500", func(t *testing.T) {
		formData := url.Values{}
		formData.Set("date", "2024-10-16")

		body := strings.NewReader(formData.Encode())

		request, err := http.NewRequestWithContext(ctx, http.MethodPost, "/items", body)
		assert.NilError(t, err)

		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		recorder := httptest.NewRecorder()

		categoryService.
			EXPECT().
			GetCategories(ctx).
			Return(models.Categories{}, nil)

		itemService.
			EXPECT().
			CreateItem(ctx, gomock.Any()).
			Return(errors.New("create item error"))

		handler.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})
}
