//nolint:dupl
package helpers_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/tksasha/balance/internal/core/common/helpers"
	"gotest.tools/v3/assert"
)

func TestParseYear(t *testing.T) {
	t.Run("returns primary value if it is not zero", func(t *testing.T) {
		actual := helpers.Year(2025, "2024")

		expected := "2025"

		assert.Equal(t, actual, expected)
	})

	t.Run("returns parsed value if it is parsed", func(t *testing.T) {
		actual := helpers.Year(0, "2025")

		expected := "2025"

		assert.Equal(t, actual, expected)
	})

	t.Run("returns current year if parsing failed", func(t *testing.T) {
		actual := helpers.Year(0, "abc")

		expected := fmt.Sprintf(
			"%04d",
			time.Now().Year(),
		)

		assert.Equal(t, actual, expected)
	})

	t.Run("returns current year if parsed value is zero", func(t *testing.T) {
		actual := helpers.Year(0, "0")

		expected := fmt.Sprintf(
			"%04d",
			time.Now().Year(),
		)

		assert.Equal(t, actual, expected)
	})
}
