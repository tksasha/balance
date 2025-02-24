package components

import (
	"github.com/tksasha/balance/internal/app/category"
	"github.com/tksasha/validator"
	. "maragu.dev/gomponents"      //nolint: stylecheck
	. "maragu.dev/gomponents/html" //nolint: stylecheck
)

func (c *CategoryComponent) Update(category *category.Category, errors validator.Errors) Node {
	return Div(
		c.form(category, errors),
	)
}
