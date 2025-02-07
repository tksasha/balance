package handlers_test

import (
	"context"
	"testing"

	"github.com/tksasha/balance/pkg/currencies"
)

func usdContext(ctx context.Context, t *testing.T) context.Context {
	t.Helper()

	return context.WithValue(ctx, currencies.CurrencyContextValue{}, currencies.USD)
}

func eurContext(ctx context.Context, t *testing.T) context.Context {
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
