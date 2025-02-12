package utils

import (
	"context"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/tksasha/balance/internal/middlewares"
)

type Params map[string]string

func NewMux(t *testing.T, pattern string, handler http.Handler) *http.ServeMux {
	t.Helper()

	next := handler

	middlewares := middlewares.New()

	for _, middleware := range middlewares {
		next = middleware.Wrap(next)
	}

	mux := http.NewServeMux()
	mux.Handle(pattern, next)

	return mux
}

func newRequest(ctx context.Context, t *testing.T, method, endpoint string, params Params) *http.Request {
	t.Helper()

	formData := url.Values{}

	for name, value := range params {
		formData.Add(name, value)
	}

	body := strings.NewReader(formData.Encode())

	request, err := http.NewRequestWithContext(ctx, method, endpoint, body)
	if err != nil {
		t.Fatalf("failed to build new request with context, error: %v", err)
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return request
}

func NewGetRequest(ctx context.Context, t *testing.T, endpoint string) *http.Request {
	t.Helper()

	return newRequest(ctx, t, http.MethodGet, endpoint, nil)
}
