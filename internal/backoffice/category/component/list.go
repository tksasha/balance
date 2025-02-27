package component

import (
	"github.com/tksasha/balance/internal/backoffice/category"
	. "maragu.dev/gomponents"      //nolint: stylecheck
	. "maragu.dev/gomponents/html" //nolint: stylecheck
)

func (c *Component) List(categories category.Categories) Node {
	return Map(categories, c.Category)
}

func (c *Component) Category(category *category.Category) Node {
	return Div(
		Text(category.Name),
	)
}
