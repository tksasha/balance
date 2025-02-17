package components

import (
	"github.com/tksasha/balance/internal/core/cash"
	"github.com/tksasha/balance/internal/core/common/components"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func Index(cashes cash.Cashes) Node {
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
						Text(components.Money(cash.Sum)),
					),
				)
			}),
		),
	)
}
