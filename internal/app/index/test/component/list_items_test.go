package component_test

import (
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/tksasha/balance/internal/app/index/component"
	"gotest.tools/v3/assert"
)

func TestListItems(t *testing.T) {
	component := component.New()

	t.Run("renders passed values if it is not zero", func(t *testing.T) {
		values := url.Values{"month": {"12"}, "year": {"2025"}}

		actual := component.ListItems(2007, 10, values)

		expected := "/items?month=10&year=2007"

		assert.Equal(t, actual, expected)
	})

	t.Run("renders date from request if it is present", func(t *testing.T) {
		values := url.Values{"month": {"5"}, "year": {"2024"}}

		actual := component.ListItems(0, 0, values)

		expected := "/items?month=05&year=2024"

		assert.Equal(t, actual, expected)
	})

	t.Run("renders current date if request invalid", func(t *testing.T) {
		values := url.Values{"month": {"abc"}, "year": {"abc"}}

		actual := component.ListItems(0, 0, values)

		month, year := time.Now().Month(), time.Now().Year()

		expected := fmt.Sprintf("/items?month=%02d&year=%04d", month, year)

		assert.Equal(t, actual, expected)
	})

	t.Run("renders currency if it was in query", func(t *testing.T) {
		values := url.Values{"currency": {"eur"}}

		actual := component.ListItems(2025, 12, values)

		expected := "/items?currency=eur&month=12&year=2025"

		assert.Equal(t, actual, expected)
	})
}
