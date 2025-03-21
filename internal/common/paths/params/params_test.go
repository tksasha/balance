package params_test

import (
	"net/url"
	"testing"

	"github.com/tksasha/balance/internal/common/currency"
	"github.com/tksasha/balance/internal/common/paths/params"
	"gotest.tools/v3/assert"
)

func TestNew(t *testing.T) {
	t.Run("by default", func(t *testing.T) {
		params := params.New()

		assert.Equal(t, params.Get("currency"), "uah")
	})

	t.Run("when currency is set", func(t *testing.T) {
		ds := map[string]url.Values{
			"uah": {"currency": {"uah"}},
			"usd": {"currency": {"usd"}},
			"eur": {"currency": {"eur"}},
		}

		for code, values := range ds {
			params := params.New(values)

			assert.Equal(t, params.Get("currency"), code)
		}
	})
}

func TestWithCurrency(t *testing.T) {
	original := params.New()

	for curr, code := range currency.All() {
		updated := original.WithCurrency(curr)

		assert.Equal(t, original.Get("currency"), "uah")
		assert.Equal(t, updated.Get("currency"), currency.Code(code))
	}
}

func TestString(t *testing.T) {
	params := params.New()

	assert.Equal(t, params.String(), "currency=uah")
}

func TestWithMonth(t *testing.T) {
	original := params.New()

	updated := original.WithMonth(12)

	assert.Equal(t, original.Get("month"), "")
	assert.Equal(t, updated.Get("month"), "12")
}

func TestWithYear(t *testing.T) {
	original := params.New()

	updated := original.WithYear(2025)

	assert.Equal(t, original.Get("year"), "")
	assert.Equal(t, updated.Get("year"), "2025")
}

func TestWithCategory(t *testing.T) {
	original := params.New()

	updated := original.WithCategory(17)

	assert.Equal(t, original.Get("category"), "")
	assert.Equal(t, updated.Get("category"), "17")
}
