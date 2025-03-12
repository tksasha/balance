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

	ds := []struct {
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

	for _, d := range ds {
		assert.Equal(t, path.Items(d.values, d.params), d.path)
	}
}

func TestUpdateItem(t *testing.T) {
	ds := []struct {
		values url.Values
		id     int
		path   string
	}{
		{
			values: nil,
			id:     6,
			path:   "/items/6?currency=uah",
		},
		{
			values: url.Values{"currency": {"uah"}},
			id:     6,
			path:   "/items/6?currency=uah",
		},
		{
			values: url.Values{"currency": {"usd"}},
			id:     6,
			path:   "/items/6?currency=usd",
		},
		{
			values: url.Values{"currency": {"eur"}},
			id:     6,
			path:   "/items/6?currency=eur",
		},
	}

	for _, d := range ds {
		assert.Equal(t, d.path, path.UpdateItem(d.values, d.id))
	}
}

func TestNewItem(t *testing.T) {
	ds := []struct {
		values url.Values
		path   string
	}{
		{
			values: nil,
			path:   "/items/new?currency=uah",
		},
		{
			values: url.Values{"currency": {"uah"}},
			path:   "/items/new?currency=uah",
		},
		{
			values: url.Values{"currency": {"usd"}},
			path:   "/items/new?currency=usd",
		},
		{
			values: url.Values{"currency": {"eur"}},
			path:   "/items/new?currency=eur",
		},
	}

	for _, d := range ds {
		assert.Equal(t, d.path, path.NewItem(d.values))
	}
}

func TestCreateItem(t *testing.T) {
	ds := []struct {
		values url.Values
		path   string
	}{
		{
			values: nil,
			path:   "/items?currency=uah",
		},
		{
			values: url.Values{"currency": {"uah"}},
			path:   "/items?currency=uah",
		},
		{
			values: url.Values{"currency": {"usd"}},
			path:   "/items?currency=usd",
		},
		{
			values: url.Values{"currency": {"eur"}},
			path:   "/items?currency=eur",
		},
	}

	for _, d := range ds {
		assert.Equal(t, d.path, path.CreateItem(d.values))
	}
}

func TestDeleteItem(t *testing.T) {
	ds := []struct {
		values url.Values
		id     int
		path   string
	}{
		{
			values: nil,
			id:     6,
			path:   "/items/6?currency=uah",
		},
		{
			values: url.Values{"currency": {"uah"}},
			id:     6,
			path:   "/items/6?currency=uah",
		},
		{
			values: url.Values{"currency": {"usd"}},
			id:     6,
			path:   "/items/6?currency=usd",
		},
		{
			values: url.Values{"currency": {"eur"}},
			id:     6,
			path:   "/items/6?currency=eur",
		},
	}

	for _, d := range ds {
		assert.Equal(t, d.path, path.DeleteItem(d.values, d.id))
	}
}
