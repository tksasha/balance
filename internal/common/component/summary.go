package component

import (
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) Summary(sum float64) Node {
	return Tr(Class("summary"),
		Td(Text("ВСЬОГО")),
		Td(Class("sum"),
			Text(c.Money(sum)),
		),
	)
}
