package currency_test

import (
	"testing"

	"github.com/tksasha/balance/internal/common/currency"
	"gotest.tools/v3/assert"
)

func TestGetByCode(t *testing.T) {
	data := map[string]int{
		"uah": 1,
		"UAH": 1,
		"usd": 2,
		"USD": 2,
		"eur": 3,
		"EUR": 3,
		"xxx": 0,
	}

	for code, id := range data {
		assert.Equal(t, currency.GetByCode(code), currency.Currency(id))
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
