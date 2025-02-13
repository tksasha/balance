package components

import (
	"github.com/tksasha/balance/internal/category"
	. "maragu.dev/gomponents"      //nolint: stylecheck
	. "maragu.dev/gomponents/html" //nolint: stylecheck
)

func CategoryList(categories category.Categories) Node {
	return Table(
		Class("table"),
		THead(
			Tr(
				Th(Text("Name")),
			),
		),
		TBody(
			Map(categories, CategoryListItem),
		),
	)
}
