package path_test

import (
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/tksasha/balance/internal/common/component/path"
	"gotest.tools/v3/assert"
)

func TestItems(t *testing.T) {
	today := time.Now()

	data := []struct {
		values url.Values
		params path.Params
		path   string
	}{
		{
			values: nil,
			params: nil,
			path:   fmt.Sprintf("/items?currency=uah&month=%d&year=%d", today.Month(), today.Year()),
		},
		{
			values: url.Values{"currency": {"eur"}, "month": {"1"}, "year": {"2025"}},
			params: nil,
			path:   "/items?currency=eur&month=1&year=2025",
		},
		{
			values: url.Values{"currency": {"eur"}, "month": {"1"}, "year": {"2025"}},
			params: path.Params{"currency": "usd"},
			path:   "/items?currency=usd&month=1&year=2025",
		},
		{
			values: url.Values{"currency": {"eur"}, "month": {"1"}, "year": {"2025"}},
			params: path.Params{"month": "2"},
			path:   "/items?currency=eur&month=2&year=2025",
		},
		{
			values: url.Values{"currency": {"eur"}, "month": {"1"}, "year": {"2025"}},
			params: path.Params{"year": "2022"},
			path:   "/items?currency=eur&month=1&year=2022",
		},
	}

	for _, d := range data {
		assert.Equal(t, path.Items(d.values, d.params), d.path)
	}
}
