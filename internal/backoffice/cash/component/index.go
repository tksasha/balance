package component

import (
	"github.com/tksasha/balance/internal/backoffice/cash"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) List(cashes cash.Cashes) Node {
	return Table(Class("w-100"),
		THead(
			Tr(
				Th(Text("Name")),
				Th(),
			),
		),
		TBody(
			Map(cashes, c.cash),
		),
	)
}

func (c *Component) cash(cash *cash.Cash) Node {
	return Tr(
		Td(
			Text(cash.Name),
		),
		Td(
			Div(Class("btn btn-outline-primary"), htmx.Get(editCashPath(cash.ID)),
				Text("Редагувати"),
			),
		),
	)
}
