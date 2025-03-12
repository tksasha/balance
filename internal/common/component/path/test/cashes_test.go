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

func TestEditCash(t *testing.T) {
	data := []struct {
		values url.Values
		id     int
		path   string
	}{
		{
			values: nil, id: 9, path: "/cashes/9/edit?currency=uah",
		},
		{
			values: url.Values{"currency": []string{"uah"}}, id: 9, path: "/cashes/9/edit?currency=uah",
		},
		{
			values: url.Values{"currency": []string{"usd"}}, id: 9, path: "/cashes/9/edit?currency=usd",
		},
		{
			values: url.Values{"currency": []string{"eur"}}, id: 9, path: "/cashes/9/edit?currency=eur",
		},
	}

	for _, d := range data {
		assert.Equal(t, d.path, path.EditCashPath(d.values, d.id))
	}
}

func TestUpdateCash(t *testing.T) {
	data := []struct {
		values url.Values
		id     int
		path   string
	}{
		{
			values: nil, id: 9, path: "/cashes/9?currency=uah",
		},
		{
			values: url.Values{"currency": []string{"uah"}}, id: 9, path: "/cashes/9?currency=uah",
		},
		{
			values: url.Values{"currency": []string{"usd"}}, id: 9, path: "/cashes/9?currency=usd",
		},
		{
			values: url.Values{"currency": []string{"eur"}}, id: 9, path: "/cashes/9?currency=eur",
		},
	}

	for _, d := range data {
		assert.Equal(t, d.path, path.UpdateCashPath(d.values, d.id))
	}
}
