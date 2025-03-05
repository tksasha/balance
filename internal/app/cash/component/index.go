package component

import (
	"github.com/tksasha/balance/internal/app/cash"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) Index(cashes cash.Cashes) Node {
	var nodes []Node

	nodes = append(nodes, Map(cashes, func(cash *cash.Cash) Node { return c.cash(cash) }))

	if cashes.HasMoreThanOne() {
		nodes = append(nodes, c.Summary(cashes.Sum()))
	}

	return Div(Class("col-3"),
		Div(Class("card cash"),
			Div(Class("card-body"),
				Table(Class("w-100 summarize"),
					TBody(c.Map(nodes)),
				),
			),
		),
	)
}

func (c *Component) cash(cash *cash.Cash, children ...Node) Node {
	return Tr(ID(c.cashID(cash.ID)),
		Td(Text(cash.Name)),
		Td(Class("sum"),
			Div(Class("link"),
				htmx.Get(c.cashEditPath(cash.ID)),
				htmx.Target("#modal-body"),
				Data("bs-toggle", "modal"),
				Data("bs-target", "#modal"),
				Text(c.Money(cash.Sum)),
			),
		),
		c.Map(children),
	)
}
