package service_test

import (
	"testing"

	"github.com/shopspring/decimal"
)

func dec(t *testing.T, f float64) decimal.Decimal {
	t.Helper()

	return decimal.NewFromFloat(f)
}

func eq(t *testing.T, d decimal.Decimal, f float64) bool {
	t.Helper()

	return decimal.NewFromFloat(f).Equal(d)
}
