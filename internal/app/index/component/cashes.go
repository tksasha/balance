package component

import (
	"github.com/tksasha/balance/internal/app/cash"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) cashes(cashes cash.Cashes) Node {
	return Div(Class("container-fluid"),
		Div(Class("row mt-4"),
			Div(Class("col-3"),
				Div(Class("card cash"),
					Div(Class("card-body"),
						Table(Class("w-100 summarize"),
							TBody(
								Map(cashes, c.cash),
							),
						),
					),
				),
			),
		),
	)
}

func (c *Component) cash(cash *cash.Cash) Node {
	return Tr(c.id("cash", cash.ID),
		Td(Text(cash.Name)),
		Td(Class("sum"),
			Div(Class("link"),
				htmx.Get(c.editCashPath(cash.ID)),
				htmx.Target("#modal-body"),
				Data("bs-toggle", "modal"),
				Data("bs-target", "#modal"),
				Text(c.Money(cash.Sum)),
			),
		),
	)
}
