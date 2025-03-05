package component

import (
	"github.com/tksasha/balance/internal/app/cash"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) List(cashes cash.Cashes) Node {
	var nodes []Node

	nodes = append(nodes, Map(cashes, c.cash))

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
