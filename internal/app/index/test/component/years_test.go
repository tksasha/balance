package component_test

import (
	"bytes"
	"testing"

	"github.com/tksasha/balance/internal/app/index/component"
	"github.com/tksasha/balance/internal/common/currency"
	"github.com/tksasha/balance/internal/common/paths/params"
	"gotest.tools/v3/golden"
)

func TestYears(t *testing.T) {
	component := component.New()

	testdata := map[currency.Currency]string{
		0:            "years.html",
		currency.UAH: "years-uah.html",
		currency.USD: "years-usd.html",
		currency.EUR: "years-eur.html",
	}

	for currency, filename := range testdata {
		w := bytes.NewBuffer([]byte{})

		params := params.New().WithCurrency(currency)

		if err := component.Years(params).Render(w); err != nil {
			t.Fatal(err)
		}

		golden.Assert(t, w.String(), filename)
	}
}
