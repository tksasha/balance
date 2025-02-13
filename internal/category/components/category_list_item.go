package components

import (
	"github.com/tksasha/balance/internal/category"
	. "maragu.dev/gomponents"      //nolint: stylecheck
	. "maragu.dev/gomponents/html" //nolint: stylecheck
)

func CategoryListItem(category *category.Category) Node {
	return Tr(
		Td(
			Text(category.Name),
		),
	)
}
