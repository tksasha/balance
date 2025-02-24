package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/tksasha/balance/internal/common/currency"
)

type currencyMiddleware struct{}

func (m *currencyMiddleware) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		code := strings.ToUpper(
			r.URL.Query().Get("currency"),
		)

		curr := currency.GetByCode(code)
		if curr == 0 {
			curr = currency.Default
		}

		ctx := context.WithValue(r.Context(), currency.ContextValue{}, curr)

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
