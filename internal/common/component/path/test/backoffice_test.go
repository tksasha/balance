package path_test

import (
	"testing"

	"github.com/tksasha/balance/internal/common/component/path"
	"gotest.tools/v3/assert"
)

func TestBackofficeCashes(t *testing.T) {
	ds := []struct {
		params path.Params
		path   string
	}{
		{
			params: nil,
			path:   "/backoffice/cashes?currency=uah",
		},
		{
			params: path.Params{"currency": "uah"},
			path:   "/backoffice/cashes?currency=uah",
		},
		{
			params: path.Params{"currency": "usd"},
			path:   "/backoffice/cashes?currency=usd",
		},
		{
			params: path.Params{"currency": "eur"},
			path:   "/backoffice/cashes?currency=eur",
		},
	}

	for _, d := range ds {
		assert.Equal(t, d.path, path.BackofficeCashes(d.params))
	}
}

func TestEditBackofficeCash(t *testing.T) {
	ds := []struct {
		params path.Params
		id     int
		path   string
	}{
		{
			params: nil,
			id:     1,
			path:   "/backoffice/cashes/1/edit?currency=uah",
		},
		{
			params: path.Params{"currency": "uah"},
			id:     2,
			path:   "/backoffice/cashes/2/edit?currency=uah",
		},
		{
			params: path.Params{"currency": "usd"},
			id:     3,
			path:   "/backoffice/cashes/3/edit?currency=usd",
		},
		{
			params: path.Params{"currency": "eur"},
			id:     4,
			path:   "/backoffice/cashes/4/edit?currency=eur",
		},
	}

	for _, d := range ds {
		assert.Equal(t, d.path, path.EditBackofficeCash(d.params, d.id))
	}
}

func TestDeleteBackofficeCash(t *testing.T) {
	ds := []struct {
		params path.Params
		id     int
		path   string
	}{
		{
			params: nil,
			id:     1,
			path:   "/backoffice/cashes/1?currency=uah",
		},
		{
			params: path.Params{"currency": "uah"},
			id:     2,
			path:   "/backoffice/cashes/2?currency=uah",
		},
		{
			params: path.Params{"currency": "usd"},
			id:     3,
			path:   "/backoffice/cashes/3?currency=usd",
		},
		{
			params: path.Params{"currency": "eur"},
			id:     4,
			path:   "/backoffice/cashes/4?currency=eur",
		},
	}

	for _, d := range ds {
		assert.Equal(t, d.path, path.DeleteBackofficeCash(d.params, d.id))
	}
}

func TestNewBackofficeCash(t *testing.T) {
	ds := []struct {
		params path.Params
		path   string
	}{
		{
			params: nil,
			path:   "/backoffice/cashes/new?currency=uah",
		},
		{
			params: path.Params{"currency": "uah"},
			path:   "/backoffice/cashes/new?currency=uah",
		},
		{
			params: path.Params{"currency": "usd"},
			path:   "/backoffice/cashes/new?currency=usd",
		},
		{
			params: path.Params{"currency": "eur"},
			path:   "/backoffice/cashes/new?currency=eur",
		},
	}

	for _, d := range ds {
		assert.Equal(t, d.path, path.NewBackofficeCash(d.params))
	}
}

func TestUpdateBackofficeCash(t *testing.T) {
	ds := []struct {
		params path.Params
		id     int
		path   string
	}{
		{
			params: nil,
			id:     1,
			path:   "/backoffice/cashes/1?currency=uah",
		},
		{
			params: path.Params{"currency": "uah"},
			id:     2,
			path:   "/backoffice/cashes/2?currency=uah",
		},
		{
			params: path.Params{"currency": "usd"},
			id:     3,
			path:   "/backoffice/cashes/3?currency=usd",
		},
		{
			params: path.Params{"currency": "eur"},
			id:     4,
			path:   "/backoffice/cashes/4?currency=eur",
		},
	}

	for _, d := range ds {
		assert.Equal(t, d.path, path.UpdateBackofficeCash(d.params, d.id))
	}
}

func TestNewBackofficeCategory(t *testing.T) {
	ds := []struct {
		params path.Params
		path   string
	}{
		{
			params: nil,
			path:   "/backoffice/categories/new?currency=uah",
		},
		{
			params: path.Params{"currency": "uah"},
			path:   "/backoffice/categories/new?currency=uah",
		},
		{
			params: path.Params{"currency": "usd"},
			path:   "/backoffice/categories/new?currency=usd",
		},
		{
			params: path.Params{"currency": "eur"},
			path:   "/backoffice/categories/new?currency=eur",
		},
	}

	for _, d := range ds {
		assert.Equal(t, d.path, path.NewBackofficeCategory(d.params))
	}
}

func TestEditBackofficeCategory(t *testing.T) {
	ds := []struct {
		params path.Params
		id     int
		path   string
	}{
		{
			params: nil,
			id:     1,
			path:   "/backoffice/categories/1/edit?currency=uah",
		},
		{
			params: path.Params{"currency": "uah"},
			id:     2,
			path:   "/backoffice/categories/2/edit?currency=uah",
		},
		{
			params: path.Params{"currency": "usd"},
			id:     3,
			path:   "/backoffice/categories/3/edit?currency=usd",
		},
		{
			params: path.Params{"currency": "eur"},
			id:     4,
			path:   "/backoffice/categories/4/edit?currency=eur",
		},
	}

	for _, d := range ds {
		assert.Equal(t, d.path, path.EditBackofficeCategory(d.params, d.id))
	}
}
