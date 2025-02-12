package components

import (
	"github.com/tksasha/balance/internal/models"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func CashList(cashes models.Cashes) Node {
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
			Map(cashes, CashListItem),
		),
	)
}
