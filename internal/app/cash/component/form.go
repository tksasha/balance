package component

import (
	"github.com/tksasha/balance/internal/app/cash"
	"github.com/tksasha/validation"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) form(node Node, cash *cash.Cash, errors validation.Errors) Node {
	return Form(
		node,
		c.Input("Назва", "name", cash.Name, errors.Get("name")),
		c.Input("Сума", "formula", cash.Formula, errors.Get("sum")),
		c.Submit(cash.ID),
	)
}
