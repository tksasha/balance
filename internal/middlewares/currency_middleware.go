package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/tksasha/balance/internal/models"
)

type CurrencyMiddleware struct{}

func NewCurrencyMiddleware() *CurrencyMiddleware {
	return &CurrencyMiddleware{}
}

func (m *CurrencyMiddleware) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		code := strings.ToUpper(
			r.URL.Query().Get("currency"),
		)

		currency := models.GetCurrencyByCode(code)
		if currency == 0 {
			_, code = models.GetDefaultCurrency()
		}

		ctx := context.WithValue(r.Context(), models.CurrencyContextValue{}, code)

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
