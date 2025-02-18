//nolint:dupl
package valueobjects_test

import (
	"testing"

	"github.com/tksasha/balance/internal/core/common/valueobjects"
	"github.com/tksasha/balance/internal/core/common/valueobjects/mocks"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestParseMonth(t *testing.T) {
	controller := gomock.NewController(t)

	currentDateProvider := mocks.NewMockCurrentDateProvider(controller)

	month := valueobjects.NewMonth(currentDateProvider)

	t.Run("returns primary value if it is not zero", func(t *testing.T) {
		assert.Equal(t, month.Parse(3, "2"), 3)
	})

	t.Run("returns parsed value if it is parsed", func(t *testing.T) {
		assert.Equal(t, month.Parse(0, "3"), 3)
	})

	t.Run("returns current month if parsing failed", func(t *testing.T) {
		currentDateProvider.EXPECT().CurrentMonth().Return(3)

		assert.Equal(t, month.Parse(0, "abc"), 3)
	})

	t.Run("returns current month if parsed value is zero", func(t *testing.T) {
		currentDateProvider.EXPECT().CurrentMonth().Return(3)

		assert.Equal(t, month.Parse(0, "0"), 3)
	})
}
