package components

import (
	"github.com/tksasha/balance/internal/core/category"
	. "maragu.dev/gomponents"      //nolint: stylecheck
	. "maragu.dev/gomponents/html" //nolint: stylecheck
)

func (c *CategoryComponent) List(categories category.Categories) Node {
	return Table(
		Class("table"),
		THead(
			Tr(
				Th(Text("Name")),
			),
		),
		TBody(
			Map(categories, c.row),
		),
	)
}

func (c *CategoryComponent) row(category *category.Category) Node {
	return Tr(
		Td(
			Text(category.Name),
		),
	)
}
