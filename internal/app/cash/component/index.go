package component

import (
	"github.com/tksasha/balance/internal/app/cash"
	"github.com/tksasha/balance/internal/common/paths"
	"github.com/tksasha/balance/internal/common/paths/params"
	. "maragu.dev/gomponents" //nolint:staticcheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint:staticcheck
)

func (c *Component) Index(params params.Params, groupedCashes cash.GroupedCashes) Node {
	groupedCashNodes := []Node{}

	for _, key := range groupedCashes.Keys() {
		cashes := groupedCashes[key]

		var nodes []Node

		nodes = append(nodes, Map(cashes, func(cash *cash.Cash) Node { return c.cash(params, cash) }))

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

func (c *Component) cash(params params.Params, cash *cash.Cash, children ...Node) Node {
	return Tr(ID(c.cashID(cash.ID)),
		Td(Text(cash.Name)),
		Td(Class("sum"),
			Div(Class("link"),
				htmx.Get(paths.EditCash(params, cash.ID)),
				htmx.Target("#modal-body"),
				Data("bs-toggle", "modal"),
				Data("bs-target", "#modal"),
				Text(c.Money(cash.Sum)),
			),
		),
		Group(children),
	)
}
