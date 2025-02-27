package component

import (
	"github.com/tksasha/balance/internal/app/cash"
	"github.com/tksasha/validation"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) Edit(cash *cash.Cash, errors validation.Errors) Node {
	return Div(
		c.form(cash, errors),
	)
}
