package component_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/tksasha/balance/internal/core/common/component"
	"gotest.tools/v3/assert"
)

func TestListItems(t *testing.T) {
	component := component.New()

	t.Run("renders passed values if it is not zero", func(t *testing.T) {
		req := req(t, "/?month=12&year=2000")

		actual := component.ListItems(2007, 10, req)

		expected := "/items?month=10&year=2007"

		assert.Equal(t, actual, expected)
	})

	t.Run("renders date from request if it is present", func(t *testing.T) {
		req := req(t, "/?month=5&year=2024")

		actual := component.ListItems(0, 0, req)

		expected := "/items?month=05&year=2024"

		assert.Equal(t, actual, expected)
	})

	t.Run("renders current date if request invalid", func(t *testing.T) {
		req := req(t, "/?month=abc&year=abc")

		actual := component.ListItems(0, 0, req)

		expected := fmt.Sprintf("/items?month=%02d&year=%04d", time.Now().Month(), time.Now().Year())

		assert.Equal(t, actual, expected)
	})
}

func req(t *testing.T, url string) *http.Request {
	t.Helper()

	req, err := http.NewRequestWithContext(t.Context(), http.MethodGet, url, nil)
	if err != nil {
		t.Fatal(err)
	}

	return req
}
