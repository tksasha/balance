package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/tksasha/balance/pkg/currencies"
)

type currencyMiddleware struct{}

func newCurrencyMiddleware() *currencyMiddleware {
	return &currencyMiddleware{}
}

func (m *currencyMiddleware) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		code := strings.ToUpper(
			r.URL.Query().Get("currency"),
		)

		currency := currencies.GetCurrencyByCode(code)
		if currency == 0 {
			currency = currencies.DefaultCurrency
		}

		ctx := context.WithValue(r.Context(), currencies.CurrencyContextValue{}, currency)

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
