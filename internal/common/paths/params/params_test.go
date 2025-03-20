package params_test

import (
	"net/url"
	"strings"
	"testing"

	"github.com/tksasha/balance/internal/common/currency"
	"github.com/tksasha/balance/internal/common/paths/params"
	"gotest.tools/v3/assert"
)

func TestNew(t *testing.T) {
	t.Run("by default", func(t *testing.T) {
		params := params.New()

		assert.Equal(t, "uah", params.Get("currency"))
	})

	t.Run("when currency is set", func(t *testing.T) {
		ds := map[string]url.Values{
			"uah": {"currency": {"uah"}},
			"usd": {"currency": {"usd"}},
			"eur": {"currency": {"eur"}},
		}

		for code, values := range ds {
			params := params.New(values)

			assert.Equal(t, code, params.Get("currency"))
		}
	})
}

func TestSetCurrency(t *testing.T) {
	params := params.New()

	for currency, code := range currency.All() {
		params.SetCurrency(currency)

		assert.Equal(t, strings.ToLower(code), params.Get("currency"))
	}
}

func TestSetCurrencyCode(t *testing.T) {
	params := params.New()

	for _, code := range currency.All() {
		params.SetCurrencyCode(code)

		assert.Equal(t, currency.Code(code), params.Get("currency"))
	}

	t.Run("set default currency if it not found or invalid", func(t *testing.T) {
		for _, code := range []string{"", "abc"} {
			params.SetCurrencyCode(code)

			assert.Equal(t, currency.Code("uah"), params.Get("currency"))
		}
	})
}

func TestString(t *testing.T) {
	params := params.New()

	assert.Equal(t, "currency=uah", params.String())
}
