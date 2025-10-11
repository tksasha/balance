package component

import (
	"github.com/tksasha/balance/internal/backoffice/category"
	"github.com/tksasha/validation"
	. "maragu.dev/gomponents" //nolint: staticcheck
)

func (c *Component) Update(category *category.Category, errors validation.Errors) Node {
	return c.form(category, errors)
}
