package component

import (
	"github.com/tksasha/balance/internal/backoffice/category"
	"github.com/tksasha/validation"
	. "maragu.dev/gomponents" //nolint: stylecheck
)

func (c *Component) Create(category *category.Category, errors validation.Errors) Node {
	return c.form(category, errors)
}
