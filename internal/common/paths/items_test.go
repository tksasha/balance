package paths_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/tksasha/balance/internal/common/paths"
	"github.com/tksasha/balance/internal/common/paths/params"
	"gotest.tools/v3/assert"
)

func TestItems(t *testing.T) {
	month, year := time.Now().Month(), time.Now().Year()

	ds := []struct {
		path   string
		params params.Params
	}{
		{
			path:   fmt.Sprintf("/items?currency=uah&month=%d&year=%d", month, year),
			params: params.New(),
		},
		{
			path:   fmt.Sprintf("/items?currency=eur&month=%d&year=%d", month, year),
			params: params.New().SetCurrencyCode("eur"),
		},
		{
			path:   fmt.Sprintf("/items?currency=usd&month=1&year=%d", year),
			params: params.New().SetCurrencyCode("usd").SetMonth(1),
		},
		{
			path:   "/items?currency=eur&month=12&year=2022",
			params: params.New().SetCurrencyCode("eur").SetMonth(12).SetYear(2022),
		},
		{
			path:   "/items?category=16&currency=eur&month=12&year=2022",
			params: params.New().SetCurrencyCode("eur").SetMonth(12).SetYear(2022).SetCategory(16),
		},
	}

	for _, d := range ds {
		assert.Equal(t, paths.Items(d.params), d.path)
	}
}

func TestUpdateItem(t *testing.T) {
	ds := []struct {
		id     int
		path   string
		params params.Params
	}{
		{id: 6, path: "/items/6?currency=uah", params: params.New()},
		{id: 6, path: "/items/6?currency=uah", params: params.New().SetCurrencyCode("uah")},
		{id: 6, path: "/items/6?currency=usd", params: params.New().SetCurrencyCode("usd")},
		{id: 6, path: "/items/6?currency=eur", params: params.New().SetCurrencyCode("eur")},
	}

	for _, d := range ds {
		assert.Equal(t, d.path, paths.UpdateItem(d.params, d.id))
	}
}

func TestNewItem(t *testing.T) {
	ds := []struct {
		path   string
		params params.Params
	}{
		{path: "/items/new?currency=uah", params: params.New()},
		{path: "/items/new?currency=uah", params: params.New().SetCurrencyCode("uah")},
		{path: "/items/new?currency=usd", params: params.New().SetCurrencyCode("usd")},
		{path: "/items/new?currency=eur", params: params.New().SetCurrencyCode("eur")},
	}

	for _, d := range ds {
		assert.Equal(t, d.path, paths.NewItem(d.params))
	}
}

func TestCreateItem(t *testing.T) {
	ds := []struct {
		path   string
		params params.Params
	}{
		{path: "/items?currency=uah", params: params.New()},
		{path: "/items?currency=uah", params: params.New().SetCurrencyCode("uah")},
		{path: "/items?currency=usd", params: params.New().SetCurrencyCode("usd")},
		{path: "/items?currency=eur", params: params.New().SetCurrencyCode("eur")},
	}

	for _, d := range ds {
		assert.Equal(t, d.path, paths.CreateItem(d.params))
	}
}

func TestDeleteItem(t *testing.T) {
	ds := []struct {
		id     int
		path   string
		params params.Params
	}{
		{id: 6, path: "/items/6?currency=uah", params: params.New()},
		{id: 6, path: "/items/6?currency=uah", params: params.New().SetCurrencyCode("uah")},
		{id: 6, path: "/items/6?currency=usd", params: params.New().SetCurrencyCode("usd")},
		{id: 6, path: "/items/6?currency=eur", params: params.New().SetCurrencyCode("eur")},
	}

	for _, d := range ds {
		assert.Equal(t, d.path, paths.DeleteItem(d.params, d.id))
	}
}

func TestEditItem(t *testing.T) {
	ds := []struct {
		id     int
		path   string
		params params.Params
	}{
		{id: 6, path: "/items/6/edit?currency=uah", params: params.New()},
		{id: 6, path: "/items/6/edit?currency=uah", params: params.New().SetCurrencyCode("uah")},
		{id: 6, path: "/items/6/edit?currency=usd", params: params.New().SetCurrencyCode("usd")},
		{id: 6, path: "/items/6/edit?currency=eur", params: params.New().SetCurrencyCode("eur")},
	}

	for _, d := range ds {
		assert.Equal(t, d.path, paths.EditItem(d.params, d.id))
	}
}
