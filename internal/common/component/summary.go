package component

import (
	. "maragu.dev/gomponents"      //nolint:staticcheck
	. "maragu.dev/gomponents/html" //nolint:staticcheck
)

func (c *Component) Summary(sum float64) Node {
	return Tr(Class("summary"),
		Td(Text("ВСЬОГО")),
		Td(Class("sum"),
			Text(c.Money(sum)),
		),
	)
}
