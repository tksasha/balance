package middlewares

import (
	"net/http"

	"github.com/tksasha/balance/internal/common"
)

type initMiddleware struct{}

func (m *initMiddleware) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrapper := common.NewResponseWriterWrapper(w)

		next.ServeHTTP(wrapper, r)
	})
}
