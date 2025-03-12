package path_test

import (
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/tksasha/balance/internal/common/component/path"
	"gotest.tools/v3/assert"
)

func TestCategories(t *testing.T) {
	today := time.Now()

	ds := []struct {
		values url.Values
		params path.Params
		path   string
	}{
		{
			values: nil,
			params: nil,
			path:   fmt.Sprintf("/categories?currency=uah&month=%d&year=%d", today.Month(), today.Year()),
		},
		{
			values: nil,
			params: path.Params{"month": "12"},
			path:   fmt.Sprintf("/categories?currency=uah&month=12&year=%d", today.Year()),
		},
		{
			values: nil,
			params: path.Params{"year": "2022"},
			path:   fmt.Sprintf("/categories?currency=uah&month=%d&year=2022", today.Month()),
		},
		{
			values: nil,
			params: path.Params{"month": "12", "year": "2025"},
			path:   "/categories?currency=uah&month=12&year=2025",
		},
		{
			values: url.Values{"currency": []string{"uah"}},
			params: path.Params{"month": "12", "year": "2025"},
			path:   "/categories?currency=uah&month=12&year=2025",
		},
		{
			values: url.Values{"currency": []string{"usd"}},
			params: path.Params{"month": "12", "year": "2025"},
			path:   "/categories?currency=usd&month=12&year=2025",
		},
		{
			values: url.Values{"currency": []string{"eur"}},
			params: path.Params{"month": "12", "year": "2025"},
			path:   "/categories?currency=eur&month=12&year=2025",
		},
	}

	for _, d := range ds {
		assert.Equal(t, path.Categories(d.values, d.params), d.path)
	}
}
