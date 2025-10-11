package component

import (
	"maps"
	"slices"

	"github.com/tksasha/balance/internal/common/currency"
	"github.com/tksasha/balance/internal/common/paths/params"
	. "maragu.dev/gomponents"      //nolint:staticcheck
	. "maragu.dev/gomponents/html" //nolint:staticcheck
)

func (c *Component) currenciesWidget(params params.Params) Node {
	currencyToDisplay := params.Get("currency")

	if currencyToDisplay == "" {
		currencyToDisplay = currency.GetCode(currency.Default)
	}

	return Button(Type("button"), ID("currencies-widget"), Text(currency.Code(currencyToDisplay)))
}

func (c *Component) currenciesWidgetContent(params params.Params) Node {
	currentCurrency := params.Get("currency")

	if currentCurrency == "" {
		currentCurrency = currency.GetCode(currency.Default)
	}

	currentCurrency = currency.Code(currentCurrency)

	currencies := slices.Collect(maps.Values(currency.All()))

	return Div(ID("currencies-widget-content"),
		Div(Class("container"),
			Div(Class("row"),
				Map(currencies, func(curr string) Node {
					curr = currency.Code(curr)

					if curr == currentCurrency {
						return nil
					}

					return Div(Class("col"),
						Div(Class("text-center text-uppercase"),
							A(Class("currencies-widget-currency"), Href("/?currency="+curr), Text(curr)),
						),
					)
				}),
			),
		),
	)
}
