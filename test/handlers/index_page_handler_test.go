package handlers_test

import (
	"context"
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

	handler := handlers.NewIndexPageHandler(categoryService)

	request, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/", nil)
	assert.NilError(t, err)

	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}
