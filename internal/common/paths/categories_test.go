package paths_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/tksasha/balance/internal/common/currency"
	"github.com/tksasha/balance/internal/common/paths"
	"github.com/tksasha/balance/internal/common/paths/params"
	"gotest.tools/v3/assert"
)

func TestCategories(t *testing.T) {
	month, year := time.Now().Month(), time.Now().Year()

	ds := []struct {
		path   string
		params params.Params
	}{
		{
			path:   fmt.Sprintf("/categories?currency=uah&month=%d&year=%d", month, year),
			params: params.New(),
		},
		{
			path:   fmt.Sprintf("/categories?currency=uah&month=12&year=%d", year),
			params: params.New().WithMonth(12),
		},
		{
			path:   fmt.Sprintf("/categories?currency=uah&month=%d&year=2022", month),
			params: params.New().WithYear(2022),
		},
		{
			path:   "/categories?currency=uah&month=12&year=2025",
			params: params.New().WithMonth(12).WithYear(2025),
		},
		{
			path:   "/categories?currency=eur&month=12&year=2025",
			params: params.New().WithMonth(12).WithYear(2025).WithCurrency(currency.EUR),
		},
	}

	for _, d := range ds {
		assert.Equal(t, paths.Categories(d.params), d.path)
	}
}
