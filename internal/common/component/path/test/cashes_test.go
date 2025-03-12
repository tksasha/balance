package path_test

import (
	"net/url"
	"testing"

	"github.com/tksasha/balance/internal/common/component/path"
	"gotest.tools/v3/assert"
)

func TestCashes(t *testing.T) {
	data := []struct {
		path   string
		values url.Values
	}{
		{path: "/cashes?currency=uah", values: nil},
		{path: "/cashes?currency=uah", values: url.Values{"currency": []string{"uah"}}},
		{path: "/cashes?currency=usd", values: url.Values{"currency": []string{"usd"}}},
		{path: "/cashes?currency=eur", values: url.Values{"currency": []string{"eur"}}},
	}

	for _, d := range data {
		assert.Equal(t, d.path, path.Cashes(d.values))
	}
}
