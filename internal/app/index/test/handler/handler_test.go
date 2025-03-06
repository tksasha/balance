package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/app/index/handler"
	"github.com/tksasha/balance/internal/server/middlewares"
	"gotest.tools/v3/assert"
)

func TestIndexHandler(t *testing.T) {
	indexHandler := handler.New()

	mux := http.NewServeMux()

	next := http.Handler(indexHandler)

	for _, middleware := range middlewares.New() {
		next = middleware.Wrap(next)
	}

	mux.Handle("/", next)

	ctx := t.Context()

	t.Run("renders 200 when there no errors", func(t *testing.T) {
		request, err := http.NewRequestWithContext(ctx, http.MethodGet, "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)
	})
}
