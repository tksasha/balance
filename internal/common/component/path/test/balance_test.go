package path_test

import (
	"net/url"
	"testing"

	"github.com/tksasha/balance/internal/common/component/path"
	"gotest.tools/v3/assert"
)

func TestBalance(t *testing.T) {
	data := []struct {
		path   string
		values url.Values
	}{
		{path: "/balance?currency=uah", values: url.Values{}},
		{path: "/balance?currency=uah", values: url.Values{"currency": []string{"uah"}}},
		{path: "/balance?currency=usd", values: url.Values{"currency": []string{"usd"}}},
		{path: "/balance?currency=eur", values: url.Values{"currency": []string{"eur"}}},
	}

	for _, d := range data {
		assert.Equal(t, d.path, path.Balance(d.values))
	}
}
