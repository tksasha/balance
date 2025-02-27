package component

import (
	"github.com/tksasha/balance/internal/app/cash"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) List(cashes cash.Cashes) Node {
	return Map(cashes, c.cash)
}

func (c *Component) cash(cash *cash.Cash) Node {
	return c.Template(
		Tr(
			Td(Text(cash.Name)),
			Td(
				Div(
					htmx.Get(c.editPath(cash.ID)),
					Text(c.Money(cash.Sum)),
				),
			),
		),
	)
}
