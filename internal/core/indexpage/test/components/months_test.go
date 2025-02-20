package components_test

import (
	"bytes"
	"testing"

	"github.com/tksasha/balance/internal/core/common/component"
	"github.com/tksasha/balance/internal/core/common/tests"
	"github.com/tksasha/balance/internal/core/indexpage/components"
	"gotest.tools/v3/golden"
)

func TestMonths(t *testing.T) {
	component := components.NewMonthsComponent(component.New())

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
