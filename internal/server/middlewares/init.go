package middlewares

import (
	"net/http"

	"github.com/tksasha/balance/internal/core/common/response"
)

type initMiddleware struct{}

func (m *initMiddleware) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := &response.Response{ResponseWriter: w}

		next.ServeHTTP(response, r)
	})
}
