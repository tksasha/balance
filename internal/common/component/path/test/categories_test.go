package path_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/tksasha/balance/internal/common/component/path"
	"gotest.tools/v3/assert"
)

func TestCategories(t *testing.T) {
	today := time.Now()

	data := []struct {
		params path.Params
		path   string
	}{
		{nil, fmt.Sprintf("/categories?month=%d&year=%d", today.Month(), today.Year())},
		{path.Params{"month": "12"}, fmt.Sprintf("/categories?month=12&year=%d", today.Year())},
		{path.Params{"year": "2022"}, fmt.Sprintf("/categories?month=%d&year=2022", today.Month())},
		{path.Params{"month": "12", "year": "2025"}, "/categories?month=12&year=2025"},
	}

	for _, d := range data {
		assert.Equal(t, path.Categories(d.params), d.path)
	}
}
