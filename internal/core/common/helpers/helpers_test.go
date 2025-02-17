package helpers_test

import (
	"net/http"
	"testing"

	"github.com/tksasha/balance/internal/core/common/helpers"
	"github.com/tksasha/balance/internal/core/common/tests"
	"gotest.tools/v3/assert"
)

func TestItemsPath(t *testing.T) {
	ctx := t.Context()

	data := map[*http.Request]string{
		nil:                              "/items",
		tests.NewGetRequest(ctx, t, "/"): "/items",
		tests.NewGetRequest(ctx, t, "/?currency=uah"):       "/items?currency=uah",
		tests.NewGetRequest(ctx, t, "/?currency=usd"):       "/items?currency=usd",
		tests.NewGetRequest(ctx, t, "/?currency=eur"):       "/items?currency=eur",
		tests.NewGetRequest(ctx, t, "/?year=2025"):          "/items?year=2025",
		tests.NewGetRequest(ctx, t, "/?month=02"):           "/items?month=02",
		tests.NewGetRequest(ctx, t, "/?year=2025&month=02"): "/items?month=02&year=2025",
	}

	for request, expectation := range data {
		path := helpers.ItemsPath(request, 0, 0)

		assert.Equal(t, path, expectation)
	}

	t.Run("returns year when year is set", func(t *testing.T) {
		request := tests.NewGetRequest(ctx, t, "/?year=2024&month=02")

		path := helpers.ItemsPath(request, 2025, 0)

		assert.Equal(t, path, "/items?month=02&year=2025")
	})
}

func TestEditItemPath(t *testing.T) {
	data := map[*http.Request]string{
		nil: "/items/1418/edit",
	}

	for _, expectation := range data {
		path := helpers.EditItemPath(1418)

		assert.Equal(t, path, expectation)
	}
}
