package components_test

import (
	"bytes"
	"net/url"
	"testing"

	"github.com/tksasha/balance/internal/common/component"
	"github.com/tksasha/balance/internal/core/index/components"
	"gotest.tools/v3/golden"
)

func TestYears(t *testing.T) {
	component := components.NewYearsComponent(component.New())

	testdata := map[string]string{
		"":    "years.html",
		"uah": "years-uah.html",
		"usd": "years-usd.html",
		"eur": "years-eur.html",
	}

	for currency, filename := range testdata {
		w := bytes.NewBuffer([]byte{})

		values := url.Values{}

		if currency != "" {
			values.Add("currency", currency)
		}

		if err := component.Years(values).Render(w); err != nil {
			t.Fatal(err)
		}

		golden.Assert(t, w.String(), filename)
	}
}
