package component

import (
	"github.com/tksasha/balance/internal/app/cash"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) List(cashes cash.Cashes) Node {
	return Div(Class("col-3"),
		Div(Class("card cash"),
			Div(Class("card-body"),
				Table(Class("w-100 summarize"),
					TBody(Map(cashes, c.cash)),
				),
			),
		),
	)
}
