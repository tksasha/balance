package middlewares

import (
	"log/slog"
	"net/http"

	"github.com/tksasha/balance/internal/core/common/response"
)

type logMiddleware struct{}

func (m *logMiddleware) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, ok := w.(*response.Response)
		if !ok {
			slog.Error("invalid response writer", "response", response)

			return
		}

		next.ServeHTTP(response, r)

		code := http.StatusOK

		if response.Code != 0 {
			code = response.Code
		}

		slog.Info(
			"request",
			"method", r.Method,
			"url", r.URL.String(),
			"status", code,
		)
	})
}
