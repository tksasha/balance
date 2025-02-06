package testutils

import (
	"context"
	"testing"

	"github.com/tksasha/balance/pkg/currencies"
)

func USDContext(ctx context.Context, t *testing.T) context.Context {
	t.Helper()

	return context.WithValue(ctx, currencies.CurrencyContextValue{}, currencies.USD)
}

func EURContext(ctx context.Context, t *testing.T) context.Context {
	t.Helper()

	return context.WithValue(ctx, currencies.CurrencyContextValue{}, currencies.EUR)
}

func getCurrency(ctx context.Context) currencies.Currency {
	currency, ok := ctx.Value(currencies.CurrencyContextValue{}).(currencies.Currency)
	if !ok {
		currency = currencies.DefaultCurrency
	}

	return currency
}
