package component_test

import (
	"bytes"
	"testing"

	"github.com/tksasha/balance/internal/app/index/component"
	"github.com/tksasha/balance/internal/common/currency"
	"github.com/tksasha/balance/internal/common/paths/params"
	"gotest.tools/v3/golden"
)

func TestMonths(t *testing.T) {
	component := component.New()

	testdata := map[currency.Currency]string{
		0:            "months.html",
		currency.UAH: "months-uah.html",
		currency.USD: "months-usd.html",
		currency.EUR: "months-eur.html",
	}

	for currency, filename := range testdata {
		writer := bytes.NewBuffer([]byte{})

		params := params.New().WithCurrency(currency)

		if err := component.Months(params).Render(writer); err != nil {
			t.Fatalf("failed to render months: %v", err)
		}

		golden.Assert(t, writer.String(), filename)
	}
}
