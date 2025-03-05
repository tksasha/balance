package component

import (
	"github.com/shopspring/decimal"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) Summary(sum decimal.Decimal) Node {
	return Tr(Class("summary"),
		Td(Text("ВСЬОГО")),
		Td(Class("sum"),
			Text(c.Money(sum)),
		),
	)
}
