package components

import (
	"github.com/tksasha/balance/internal/app/cash"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *CashComponent) List(cashes cash.Cashes) Node {
	return Table(
		THead(
			Tr(
				Text("Name"),
			),
			Tr(
				Text("Sum"),
			),
		),
		TBody(
			Map(cashes, func(cash *cash.Cash) Node {
				return Tr(
					Td(
						Text(cash.Name),
					),
					Td(
						Text(c.Money(cash.Sum)),
					),
				)
			}),
		),
	)
}
