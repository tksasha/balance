package middlewares

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/tksasha/balance/internal/common/middleware"
)

type logMiddleware struct{}

func (m *logMiddleware) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrapper, ok := w.(*middleware.ResponseWriterWrapper)
		if !ok {
			slog.Error("failed to assert wrapper", "wrapper", wrapper)

			return
		}

		start := time.Now()

		next.ServeHTTP(wrapper, r)

		slog.Info(
			"request",
			"method", r.Method,
			"url", r.URL.String(),
			"status", wrapper.Code,
			"duration", time.Since(start),
		)
	})
}
