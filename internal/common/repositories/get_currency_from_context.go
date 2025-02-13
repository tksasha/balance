package repositories

import (
	"context"

	"github.com/tksasha/balance/pkg/currencies"
)

func GetCurrencyFromContext(ctx context.Context) currencies.Currency {
	currency, ok := ctx.Value(currencies.CurrencyContextValue{}).(currencies.Currency)
	if !ok {
		currency = currencies.DefaultCurrency
	}

	return currency
}
