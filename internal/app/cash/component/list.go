package component

import (
	"github.com/tksasha/balance/internal/app/cash"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *CashComponent) List(cashes cash.Cashes) Node {
	return Map(cashes, c.Cash)
}

func (c *CashComponent) Cash(cash *cash.Cash) Node {
	return c.Template(
		Tr(
			Td(Text(cash.Name)),
			Td(
				Div(
					htmx.Get(c.EditCash(cash.ID)),
					Text(c.Money(cash.Sum)),
				),
			),
		),
	)
}
