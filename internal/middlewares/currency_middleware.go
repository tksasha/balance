package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/tksasha/balance/internal/models"
)

type currencyMiddleware struct{}

func NewCurrencyMiddleware() Middleware {
	return &currencyMiddleware{}
}

func (m *currencyMiddleware) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		code := strings.ToUpper(
			r.URL.Query().Get("currency"),
		)

		currency := models.GetCurrencyByCode(code)
		if currency == 0 {
			currency, _ = models.GetDefaultCurrency()
		}

		ctx := context.WithValue(r.Context(), models.CurrencyContextValue{}, currency)

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
