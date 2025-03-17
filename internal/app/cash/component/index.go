package component

import (
	"net/url"

	"github.com/tksasha/balance/internal/app/cash"
	"github.com/tksasha/balance/internal/common/component/path"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) Index(values url.Values, groupedCashes cash.GroupedCashes) Node {
	groupedCashNodes := []Node{}

	for _, key := range groupedCashes.Keys() {
		cashes := groupedCashes[key]

		var nodes []Node

		nodes = append(nodes, Map(cashes, func(cash *cash.Cash) Node { return c.cash(values, cash) }))

		if cashes.HasMoreThanOne() {
			nodes = append(nodes, c.Summary(cashes.Sum()))
		}

		groupedCashNodes = append(groupedCashNodes, Div(Class("col-3"),
			Div(Class("card cash"),
				Div(Class("card-body"),
					Table(Class("w-100 summarize"),
						TBody(Group(nodes)),
					),
				),
			),
		))
	}

	return Group(groupedCashNodes)
}

func (c *Component) cash(values url.Values, cash *cash.Cash, children ...Node) Node {
	return Tr(ID(c.cashID(cash.ID)),
		Td(Text(cash.Name)),
		Td(Class("sum"),
			Div(Class("link"),
				htmx.Get(path.EditCash(values, cash.ID)),
				htmx.Target("#modal-body"),
				Data("bs-toggle", "modal"),
				Data("bs-target", "#modal"),
				Text(c.Money(cash.Sum)),
			),
		),
		Group(children),
	)
}
