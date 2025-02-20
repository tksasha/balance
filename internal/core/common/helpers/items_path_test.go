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

	type params struct {
		year  int
		month int
		path  string
	}

	testmap := map[*http.Request]params{
		nil: {2025, 3, "/items?month=03&year=2025"},
		nil: {0, 3, "/items?month=03&year=2025"},
		nil: {2025, 0, "/items?month=02&year=2025"},
		tests.NewGetRequest(ctx, t, "/?year=2025"): {0, 3, "/items?month=03&year=2025"},
		tests.NewGetRequest(ctx, t, "/?month=3"):   {2025, 0, "/items?month=03&year=2025"},
		tests.NewGetRequest(ctx, t, "/?year=2023"): {2025, 3, "/items?month=03&year=2025"},
		tests.NewGetRequest(ctx, t, "/?month=2"):   {2025, 3, "/items?month=03&year=2025"},
		tests.NewGetRequest(ctx, t, "/"):           {2025, 3, "/items?month=03&year=2025"},
		tests.NewGetRequest(ctx, t, "/"):           {2025, 3, "/items?month=03&year=2025"},
		tests.NewGetRequest(ctx, t, "/?year=abc"):  {0, 3, "/items?month=03&year=2025"},
		tests.NewGetRequest(ctx, t, "/?month=abc"): {2025, 0, "/items?month=02&year=2025"},
	}

	for request, params := range testmap {
		path := helpers.ItemsPath(request, params.year, params.month)

		assert.Equal(t, path, params.path)
	}
}
