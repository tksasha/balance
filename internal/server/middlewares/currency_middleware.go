package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/server/app"
)

type CurrencyMiddleware struct {
	currencies models.Currencies
	currency   models.Currency
}

func NewCurrencyMiddleware(app *app.App) *CurrencyMiddleware {
	return &CurrencyMiddleware{
		currencies: app.Currencies,
		currency:   app.Currency, // default currency
	}
}

func (m *CurrencyMiddleware) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		currencyCode := strings.ToUpper(
			r.URL.Query().Get("currency"),
		)

		currency, ok := m.currencies[currencyCode]
		if !ok {
			currency = m.currency // set default currency
		}

		ctx := context.WithValue(r.Context(), models.CurrencyContextValue{}, currency)

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
