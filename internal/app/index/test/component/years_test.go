package component_test

import (
	"bytes"
	"testing"

	"github.com/tksasha/balance/internal/app/index/component"
	"github.com/tksasha/balance/internal/common/paths/params"
	"gotest.tools/v3/golden"
)

func TestYears(t *testing.T) {
	component := component.New()

	testdata := map[string]string{
		"":    "years.html",
		"uah": "years-uah.html",
		"usd": "years-usd.html",
		"eur": "years-eur.html",
	}

	for currency, filename := range testdata {
		w := bytes.NewBuffer([]byte{})

		params := params.New().SetCurrencyCode(currency)

		if err := component.Years(params).Render(w); err != nil {
			t.Fatal(err)
		}

		golden.Assert(t, w.String(), filename)
	}
}
