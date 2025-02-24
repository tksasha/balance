package components

import (
	"github.com/tksasha/balance/internal/app/category"
	. "maragu.dev/gomponents"      //nolint: stylecheck
	. "maragu.dev/gomponents/html" //nolint: stylecheck
)

func (c *CategoryComponent) Edit(category *category.Category) Node {
	return Div(
		c.form(category, nil),
	)
}
