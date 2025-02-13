package handlers_test

import (
	"context"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/tksasha/balance/internal/middlewares"
)

type Params map[string]string

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

func newPostRequest(ctx context.Context, t *testing.T, endpoint string, params Params) *http.Request {
	t.Helper()

	return newRequest(ctx, t, http.MethodPost, endpoint, params)
}

func newPatchRequest(ctx context.Context, t *testing.T, endpoint string, params Params) *http.Request {
	t.Helper()

	return newRequest(ctx, t, http.MethodPatch, endpoint, params)
}

func newGetRequest(ctx context.Context, t *testing.T, endpoint string) *http.Request {
	t.Helper()

	return newRequest(ctx, t, http.MethodGet, endpoint, nil)
}

func newInvalidRequest(ctx context.Context, t *testing.T, method, endpoint string) *http.Request {
	t.Helper()

	body := strings.NewReader("%")

	request, err := http.NewRequestWithContext(ctx, method, endpoint, body)
	if err != nil {
		t.Fatalf("failed to build new request with context, error: %v", err)
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return request
}

func newInvalidPatchRequest(ctx context.Context, t *testing.T, endpoint string) *http.Request {
	t.Helper()

	return newInvalidRequest(ctx, t, http.MethodPatch, endpoint)
}

func newInvalidPostRequest(ctx context.Context, t *testing.T, endpoint string) *http.Request {
	t.Helper()

	return newInvalidRequest(ctx, t, http.MethodPost, endpoint)
}

func newMux(t *testing.T, pattern string, handler http.Handler) *http.ServeMux {
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
