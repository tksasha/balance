package path_test

import (
	"net/url"
	"testing"

	"github.com/tksasha/balance/internal/common/component/path"
	"gotest.tools/v3/assert"
)

func TestCashes(t *testing.T) {
	ds := []struct {
		path   string
		values url.Values
	}{
		{
			path:   "/cashes?currency=uah",
			values: nil,
		},
		{
			path:   "/cashes?currency=uah",
			values: url.Values{"currency": {"uah"}},
		},
		{
			path:   "/cashes?currency=usd",
			values: url.Values{"currency": {"usd"}},
		},
		{
			path:   "/cashes?currency=eur",
			values: url.Values{"currency": {"eur"}},
		},
	}

	for _, d := range ds {
		assert.Equal(t, d.path, path.Cashes(d.values))
	}
}

func TestEditCash(t *testing.T) {
	ds := []struct {
		values url.Values
		id     int
		path   string
	}{
		{
			values: nil,
			id:     9,
			path:   "/cashes/9/edit?currency=uah",
		},
		{
			values: url.Values{"currency": {"uah"}},
			id:     9,
			path:   "/cashes/9/edit?currency=uah",
		},
		{
			values: url.Values{"currency": {"usd"}},
			id:     9,
			path:   "/cashes/9/edit?currency=usd",
		},
		{
			values: url.Values{"currency": {"eur"}},
			id:     9,
			path:   "/cashes/9/edit?currency=eur",
		},
	}

	for _, d := range ds {
		assert.Equal(t, d.path, path.EditCash(d.values, d.id))
	}
}

func TestUpdateCash(t *testing.T) {
	ds := []struct {
		values url.Values
		id     int
		path   string
	}{
		{
			values: nil,
			id:     9,
			path:   "/cashes/9?currency=uah",
		},
		{
			values: url.Values{"currency": {"uah"}},
			id:     9,
			path:   "/cashes/9?currency=uah",
		},
		{
			values: url.Values{"currency": {"usd"}},
			id:     9,
			path:   "/cashes/9?currency=usd",
		},
		{
			values: url.Values{"currency": {"eur"}},
			id:     9,
			path:   "/cashes/9?currency=eur",
		},
	}

	for _, d := range ds {
		assert.Equal(t, d.path, path.UpdateCash(d.values, d.id))
	}
}
