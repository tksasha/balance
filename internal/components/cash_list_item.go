package components

import (
	"github.com/tksasha/balance/internal/models"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func CashListItem(cash *models.Cash) Node {
	return Tr(
		Td(
			Text(cash.Name),
		),
		Td(
			Text(
				sum(cash.Sum),
			),
		),
	)
}
