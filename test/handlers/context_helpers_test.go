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

func currencyContext(ctx context.Context, t *testing.T, currency currencies.Currency) context.Context {
	t.Helper()

	switch currency {
	case currencies.UAH:
		return context.WithValue(ctx, currencies.CurrencyContextValue{}, currencies.UAH)
	case currencies.USD:
		return context.WithValue(ctx, currencies.CurrencyContextValue{}, currencies.USD)
	case currencies.EUR:
		return context.WithValue(ctx, currencies.CurrencyContextValue{}, currencies.EUR)
	default:
		t.Fatalf("invalid currency: %v", currency)
	}

	return nil
}
