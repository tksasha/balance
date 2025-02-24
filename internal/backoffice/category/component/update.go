package component

import (
	"github.com/tksasha/balance/internal/backoffice/category"
	"github.com/tksasha/validator"
	. "maragu.dev/gomponents"      //nolint: stylecheck
	. "maragu.dev/gomponents/html" //nolint: stylecheck
)

func (c *CategoryComponent) Update(category *category.Category, errors validator.Errors) Node {
	return Div(
		c.form(category, errors),
	)
}
