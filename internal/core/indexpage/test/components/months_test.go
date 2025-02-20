package components_test

import (
	"bytes"
	"testing"

	"github.com/tksasha/balance/internal/core/common"
	"github.com/tksasha/balance/internal/core/common/helpers"
	"github.com/tksasha/balance/internal/core/common/tests"
	"github.com/tksasha/balance/internal/core/common/valueobjects/mocks"
	"github.com/tksasha/balance/internal/core/indexpage/components"
	"go.uber.org/mock/gomock"
	"gotest.tools/v3/golden"
)

func TestMonths(t *testing.T) {
	controller := gomock.NewController(t)

	currentDateProvider := mocks.NewMockCurrentDateProvider(controller)

	currentDateProvider.EXPECT().CurrentYear().Return(2025).AnyTimes()
	currentDateProvider.EXPECT().CurrentMonth().Return(3).AnyTimes()

	helpers := helpers.New(currentDateProvider)

	component := components.NewMonthsComponent(common.NewBaseComponent(helpers))

	ctx := t.Context()

	t.Run("renders months.html for empty request", func(t *testing.T) {
		writer := bytes.NewBuffer([]byte{})

		request := tests.NewGetRequest(ctx, t, "/")

		if err := component.Months(request).Render(writer); err != nil {
			t.Fatalf("failed to render months: %v", err)
		}

		golden.Assert(t, writer.String(), "months.html")
	})

	t.Run("renders months-usd.html when currency is usd", func(t *testing.T) {
		writer := bytes.NewBuffer([]byte{})

		request := tests.NewGetRequest(ctx, t, "/?currency=usd")

		if err := component.Months(request).Render(writer); err != nil {
			t.Fatalf("failed to render months: %v", err)
		}

		golden.Assert(t, writer.String(), "months-usd.html")
	})
}
