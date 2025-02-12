package components

import (
	"github.com/tksasha/balance/internal/cash"
	. "github.com/tksasha/balance/internal/components" //nolint:stylecheck
	. "maragu.dev/gomponents"                          //nolint:stylecheck
	. "maragu.dev/gomponents/html"                     //nolint:stylecheck
)

func CashListItem(cash *cash.Cash) Node {
	return Tr(
		Td(
			Text(cash.Name),
		),
		Td(
			Text(
				Sum(cash.Sum),
			),
		),
	)
}
