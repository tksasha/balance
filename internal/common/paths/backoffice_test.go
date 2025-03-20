package paths_test

import (
	"testing"

	"github.com/tksasha/balance/internal/common/paths"
	"github.com/tksasha/balance/internal/common/paths/params"
	"gotest.tools/v3/assert"
)

func TestBackofficeCashes(t *testing.T) {
	ds := []struct {
		path   string
		params params.Params
	}{
		{path: "/backoffice/cashes?currency=uah", params: params.New()},
		{path: "/backoffice/cashes?currency=uah", params: params.New().SetCurrencyCode("uah")},
		{path: "/backoffice/cashes?currency=usd", params: params.New().SetCurrencyCode("usd")},
		{path: "/backoffice/cashes?currency=eur", params: params.New().SetCurrencyCode("eur")},
	}

	for _, d := range ds {
		assert.Equal(t, d.path, paths.BackofficeCashes(d.params))
	}
}

func TestEditBackofficeCash(t *testing.T) {
	assert.Equal(t, "/backoffice/cashes/9/edit", paths.EditBackofficeCash(9))
}

func TestDeleteBackofficeCash(t *testing.T) {
	assert.Equal(t, "/backoffice/cashes/9", paths.DeleteBackofficeCash(9))
}

func TestUpdateBackofficeCash(t *testing.T) {
	assert.Equal(t, "/backoffice/cashes/9", paths.UpdateBackofficeCash(9))
}

func TestNewBackofficeCash(t *testing.T) {
	ds := []struct {
		path   string
		params params.Params
	}{
		{path: "/backoffice/cashes/new?currency=uah", params: params.New()},
		{path: "/backoffice/cashes/new?currency=uah", params: params.New().SetCurrencyCode("uah")},
		{path: "/backoffice/cashes/new?currency=usd", params: params.New().SetCurrencyCode("usd")},
		{path: "/backoffice/cashes/new?currency=eur", params: params.New().SetCurrencyCode("eur")},
	}

	for _, d := range ds {
		assert.Equal(t, d.path, paths.NewBackofficeCash(d.params))
	}
}

func TestCreateBackofficeCash(t *testing.T) {
	assert.Equal(t, "/backoffice/cashes", paths.CreateBackofficeCash())
}

func TestBackofficeCategories(t *testing.T) {
	ds := []struct {
		path   string
		params params.Params
	}{
		{path: "/backoffice/categories?currency=uah", params: params.New()},
		{path: "/backoffice/categories?currency=uah", params: params.New().SetCurrencyCode("uah")},
		{path: "/backoffice/categories?currency=usd", params: params.New().SetCurrencyCode("usd")},
		{path: "/backoffice/categories?currency=eur", params: params.New().SetCurrencyCode("eur")},
	}

	for _, d := range ds {
		assert.Equal(t, d.path, paths.BackofficeCategories(d.params))
	}
}

func TestNewBackofficeCategory(t *testing.T) {
	ds := []struct {
		path   string
		params params.Params
	}{
		{path: "/backoffice/categories/new?currency=uah", params: params.New()},
		{path: "/backoffice/categories/new?currency=uah", params: params.New().SetCurrencyCode("uah")},
		{path: "/backoffice/categories/new?currency=usd", params: params.New().SetCurrencyCode("usd")},
		{path: "/backoffice/categories/new?currency=eur", params: params.New().SetCurrencyCode("eur")},
	}

	for _, d := range ds {
		assert.Equal(t, d.path, paths.NewBackofficeCategory(d.params))
	}
}

func TestCreateBackofficeCategory(t *testing.T) {
	assert.Equal(t, "/backoffice/categories", paths.CreateBackofficeCategory())
}

func TestEditBackofficeCategory(t *testing.T) {
	assert.Equal(t, "/backoffice/categories/9/edit", paths.EditBackofficeCategory(9))
}

func TestUpdateBackofficeCategory(t *testing.T) {
	assert.Equal(t, "/backoffice/categories/9", paths.UpdateBackofficeCategory(9))
}
