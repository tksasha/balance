package component

import (
	"github.com/tksasha/balance/internal/app/cash"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) List(cashes cash.Cashes) Node {
	var nodes []Node

	nodes = append(nodes, Map(cashes, c.cash))

	if len(cashes) > 1 {
		var sum float64

		for _, cash := range cashes {
			sum += cash.Sum
		}

		nodes = append(nodes, c.Summary(sum))
	}

	return Div(Class("col-3"),
		Div(Class("card cash"),
			Div(Class("card-body"),
				Table(Class("w-100 summarize"),
					TBody(Map(nodes, func(node Node) Node { return node })),
				),
			),
		),
	)
}
