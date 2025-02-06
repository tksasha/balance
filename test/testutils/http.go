package testutils

import (
	"context"
	"io"
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

func NewPostRequest(ctx context.Context, t *testing.T, endpoint string, params Params) *http.Request {
	t.Helper()

	return newRequest(ctx, t, http.MethodPost, endpoint, params)
}

func NewPatchRequest(ctx context.Context, t *testing.T, endpoint string, params Params) *http.Request {
	t.Helper()

	return newRequest(ctx, t, http.MethodPatch, endpoint, params)
}

func NewDeleteRequest(ctx context.Context, t *testing.T, endpoint string) *http.Request {
	t.Helper()

	return newRequest(ctx, t, http.MethodDelete, endpoint, nil)
}

func NewGetRequest(ctx context.Context, t *testing.T, endpoint string) *http.Request {
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

func NewInvalidPatchRequest(ctx context.Context, t *testing.T, endpoint string) *http.Request {
	t.Helper()

	return newInvalidRequest(ctx, t, http.MethodPatch, endpoint)
}

func NewInvalidPostRequest(ctx context.Context, t *testing.T, endpoint string) *http.Request {
	t.Helper()

	return newInvalidRequest(ctx, t, http.MethodPost, endpoint)
}

func GetResponseBody(t *testing.T, reader io.Reader) string {
	t.Helper()

	body, err := io.ReadAll(reader)
	if err != nil {
		t.Fatalf("failed to parse response body: %v", err)
	}

	return string(body)
}

func NewMux(t *testing.T, pattern string, handler http.Handler) *http.ServeMux {
	t.Helper()

	handler = middlewares.NewCurrencyMiddleware().Wrap(handler)

	mux := http.NewServeMux()
	mux.Handle(pattern, handler)

	return mux
}
