package component

import (
	"github.com/tksasha/balance/internal/backoffice/category"
	"github.com/tksasha/validation"
	. "maragu.dev/gomponents"      //nolint: stylecheck
	. "maragu.dev/gomponents/html" //nolint: stylecheck
)

func (c *CategoryComponent) Update(category *category.Category, errors validation.Errors) Node {
	return Div(
		c.form(category, errors),
	)
}
