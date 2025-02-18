//nolint:dupl
package valueobjects_test

import (
	"testing"

	"github.com/tksasha/balance/internal/core/common/valueobjects"
	"github.com/tksasha/balance/internal/core/common/valueobjects/mocks"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestParseYear(t *testing.T) {
	controller := gomock.NewController(t)

	currentDateProvider := mocks.NewMockCurrentDateProvider(controller)

	year := valueobjects.NewYear(currentDateProvider)

	t.Run("returns primary value if it is not zero", func(t *testing.T) {
		assert.Equal(t, year.Parse(2025, "2024"), 2025)
	})

	t.Run("returns parsed value if it is parsed", func(t *testing.T) {
		assert.Equal(t, year.Parse(0, "2025"), 2025)
	})

	t.Run("returns current year if parsing failed", func(t *testing.T) {
		currentDateProvider.EXPECT().CurrentYear().Return(2025)

		assert.Equal(t, year.Parse(0, "abc"), 2025)
	})

	t.Run("returns current year if parsed value is zero", func(t *testing.T) {
		currentDateProvider.EXPECT().CurrentYear().Return(2025)

		assert.Equal(t, year.Parse(0, "0"), 2025)
	})
}
