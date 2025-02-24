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
