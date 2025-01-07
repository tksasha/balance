package handlers_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/handlers"
	mocksforhandlers "github.com/tksasha/balance/mocks/handlers"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestIndexPageHandler_ServeHTTP(t *testing.T) {
	controller := gomock.NewController(t)

	categoryService := mocksforhandlers.NewMockCategoryService(controller)

	ctx := context.Background()

	handler := handlers.NewIndexPageHandler(categoryService)

	t.Run("when get categories returns an error, it should respond with 500 code", func(t *testing.T) {
		request, err := http.NewRequestWithContext(ctx, http.MethodGet, "/", nil)
		assert.NilError(t, err)

		categoryService.
			EXPECT().
			GetAll(ctx).
			Return(nil, errors.New("get categories error"))

		recorder := httptest.NewRecorder()

		handler.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})
}
