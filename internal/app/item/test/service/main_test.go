package service_test

import (
	"testing"

	"github.com/shopspring/decimal"
)

func dec(t *testing.T, f float64) decimal.Decimal {
	t.Helper()

	return decimal.NewFromFloat(f)
}
