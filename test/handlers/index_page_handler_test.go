package handlers_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/handlers"
	"gotest.tools/v3/assert"
)

func TestIndexPageHandler_ServeHTTP(t *testing.T) {
	handler := handlers.NewIndexPageHandler()

	request, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/", nil)
	assert.NilError(t, err)

	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}
