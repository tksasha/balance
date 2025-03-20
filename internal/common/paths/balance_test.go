package paths_test

import (
	"testing"

	"github.com/tksasha/balance/internal/common/paths"
	"github.com/tksasha/balance/internal/common/paths/params"
	"gotest.tools/v3/assert"
)

func TestBalance(t *testing.T) {
	ds := []struct {
		path   string
		params params.Params
	}{
		{path: "/balance?currency=uah", params: params.New()},
		{path: "/balance?currency=uah", params: params.New().SetCurrencyCode("uah")},
		{path: "/balance?currency=usd", params: params.New().SetCurrencyCode("usd")},
		{path: "/balance?currency=eur", params: params.New().SetCurrencyCode("eur")},
	}

	for _, d := range ds {
		assert.Equal(t, d.path, paths.Balance(d.params))
	}
}
