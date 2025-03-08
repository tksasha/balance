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
		params path.Params
		values url.Values
		path   string
	}{
		{nil, nil, fmt.Sprintf("/items?currency=uah&month=%d&year=%d", today.Month(), today.Year())},
		{
			nil,
			url.Values{"currency": {"eur"}, "month": {"1"}, "year": {"2025"}},
			"/items?currency=eur&month=1&year=2025",
		},
		{
			map[string]string{"currency": "usd"},
			url.Values{"currency": {"eur"}, "month": {"1"}, "year": {"2025"}},
			"/items?currency=usd&month=1&year=2025",
		},
		{
			map[string]string{"month": "2"},
			url.Values{"currency": {"eur"}, "month": {"1"}, "year": {"2025"}},
			"/items?currency=eur&month=2&year=2025",
		},
		{
			map[string]string{"year": "2022"},
			url.Values{"currency": {"eur"}, "month": {"1"}, "year": {"2025"}},
			"/items?currency=eur&month=1&year=2022",
		},
	}

	for _, d := range data {
		assert.Equal(t, path.Items(d.params, d.values), d.path)
	}
}
