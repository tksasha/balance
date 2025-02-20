//nolint:dupl
package helpers_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/tksasha/balance/internal/core/common/helpers"
	"gotest.tools/v3/assert"
)

func TestMonth(t *testing.T) {
	t.Run("returns primary value if it is not zero", func(t *testing.T) {
		actual := helpers.Month(3, "2")

		expected := "03"

		assert.Equal(t, actual, expected)
	})

	t.Run("returns parsed value if it is parsed", func(t *testing.T) {
		actual := helpers.Month(0, "3")

		expected := "03"

		assert.Equal(t, actual, expected)
	})

	t.Run("returns current month if parsing failed", func(t *testing.T) {
		actual := helpers.Month(0, "abc")

		expected := fmt.Sprintf(
			"%02d",
			time.Now().Month(),
		)

		assert.Equal(t, actual, expected)
	})

	t.Run("returns current month if parsed value is zero", func(t *testing.T) {
		actual := helpers.Month(0, "0")

		expected := fmt.Sprintf(
			"%02d",
			time.Now().Month(),
		)

		assert.Equal(t, actual, expected)
	})
}
