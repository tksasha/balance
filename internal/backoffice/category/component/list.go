package component

import (
	"github.com/tksasha/balance/internal/backoffice/category"
	. "maragu.dev/gomponents"      //nolint: stylecheck
	. "maragu.dev/gomponents/html" //nolint: stylecheck
)

func (c *CategoryComponent) List(categories category.Categories) Node {
	return Map(categories, c.Category)
}

func (c *CategoryComponent) Category(category *category.Category) Node {
	return Div(
		Text(category.Name),
	)
}
