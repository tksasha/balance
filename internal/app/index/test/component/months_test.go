package component_test

import (
	"bytes"
	"testing"

	"github.com/tksasha/balance/internal/app/index/component"
	"github.com/tksasha/balance/internal/common/paths/params"
	"gotest.tools/v3/golden"
)

func TestMonths(t *testing.T) {
	component := component.New()

	testdata := map[string]string{
		"":    "months.html",
		"uah": "months-uah.html",
		"usd": "months-usd.html",
		"eur": "months-eur.html",
	}

	for currency, filename := range testdata {
		writer := bytes.NewBuffer([]byte{})

		params := params.New().SetCurrencyCode(currency)

		if err := component.Months(params).Render(writer); err != nil {
			t.Fatalf("failed to render months: %v", err)
		}

		golden.Assert(t, writer.String(), filename)
	}
}
