package components_test

import (
	"bytes"
	"net/url"
	"testing"

	"github.com/tksasha/balance/internal/core/common/component"
	"github.com/tksasha/balance/internal/core/index/components"
	"gotest.tools/v3/golden"
)

func TestMonths(t *testing.T) {
	component := components.NewMonthsComponent(component.New())

	t.Run("renders months.html for empty request", func(t *testing.T) {
		writer := bytes.NewBuffer([]byte{})

		values := url.Values{}

		if err := component.Months(values).Render(writer); err != nil {
			t.Fatalf("failed to render months: %v", err)
		}

		golden.Assert(t, writer.String(), "months.html")
	})

	t.Run("renders months-usd.html when currency is usd", func(t *testing.T) {
		writer := bytes.NewBuffer([]byte{})

		values := url.Values{"currency": {"usd"}}

		if err := component.Months(values).Render(writer); err != nil {
			t.Fatalf("failed to render months: %v", err)
		}

		golden.Assert(t, writer.String(), "months-usd.html")
	})
}
