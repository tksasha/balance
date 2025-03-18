package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/backoffice/category/handlers"
	"gotest.tools/v3/assert"
)

func TestNewHandler(t *testing.T) {
	ctx := t.Context()

	handler := handlers.NewNewHandler()

	mux := mux(t, "GET /backoffice/categories/new", handler)

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, "/backoffice/categories/new", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	mux.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}
