package component

import (
	"github.com/tksasha/balance/internal/backoffice/category"
	. "maragu.dev/gomponents"      //nolint: stylecheck
	. "maragu.dev/gomponents/html" //nolint: stylecheck
)

func (c *Component) Edit(category *category.Category) Node {
	return Div(
		c.form(category, nil),
	)
}
