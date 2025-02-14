package components

import (
	"github.com/tksasha/balance/internal/core/cash"
	"github.com/tksasha/balance/internal/core/common/components"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func CashListItem(cash *cash.Cash) Node {
	return Tr(
		Td(
			Text(cash.Name),
		),
		Td(
			Text(
				components.Sum(cash.Sum),
			),
		),
	)
}
