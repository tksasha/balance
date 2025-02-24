package components

import (
	"github.com/tksasha/balance/internal/app/category"
	"github.com/tksasha/balance/pkg/validation"
	. "maragu.dev/gomponents"      //nolint: stylecheck
	. "maragu.dev/gomponents/html" //nolint: stylecheck
)

func (c *CategoryComponent) Create(category *category.Category, errors validation.Errors) Node {
	return Div(
		c.form(category, errors),
	)
}
