package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/backoffice/index/handler"
	"gotest.tools/v3/assert"
)

func TestIndexHandler(t *testing.T) {
	handler := handler.NewIndexHandler()

	mux := http.NewServeMux()

	mux.Handle("GET /backoffice", handler)

	ctx := t.Context()

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, "/backoffice", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	mux.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)

	assert.Equal(t, "backoffice.index.shown", recorder.Header().Get("Hx-Trigger-After-Swap"))
}
