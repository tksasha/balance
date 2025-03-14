package currency_test

import (
	"maps"
	"testing"

	"github.com/tksasha/balance/internal/common/currency"
	"gotest.tools/v3/assert"
)

func TestAll(t *testing.T) {
	expected := currency.Currencies{
		currency.UAH: "UAH",
		currency.USD: "USD",
		currency.EUR: "EUR",
	}

	actual := currency.All()

	assert.Assert(t, maps.Equal(expected, actual))
}

func TestGetByCode(t *testing.T) {
	data := map[string]currency.Currency{
		"uah": currency.UAH,
		"UAH": currency.UAH,
		"usd": currency.USD,
		"USD": currency.USD,
		"eur": currency.EUR,
		"EUR": currency.EUR,
		"xxx": currency.UAH,
	}

	for code, curr := range data {
		assert.Equal(t, currency.GetByCode(code), curr)
	}
}

func TestGetDefault(t *testing.T) {
	actual := currency.Default

	expected := currency.Currency(1)

	assert.Equal(t, actual, expected)
}

func TestGetCode(t *testing.T) {
	data := []struct {
		currency currency.Currency
		code     string
	}{
		{currency.UAH, "uah"},
		{currency.USD, "usd"},
		{currency.EUR, "eur"},
		{0, "uah"},
	}

	for _, d := range data {
		assert.Equal(t, currency.GetCode(d.currency), d.code)
	}
}
