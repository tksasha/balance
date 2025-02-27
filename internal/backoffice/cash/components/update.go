package components

import (
	"github.com/tksasha/balance/internal/backoffice/cash"
	"github.com/tksasha/validation"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *CashComponent) Update(cash *cash.Cash, errors validation.Errors) Node {
	return Div(
		c.form(cash, errors),
	)
}
