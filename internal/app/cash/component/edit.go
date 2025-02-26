package component

import (
	"github.com/tksasha/balance/internal/app/cash"
	"github.com/tksasha/validator"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *CashComponent) Edit(cash *cash.Cash, errors validator.Errors) Node {
	return Div(
		c.form(cash, errors),
	)
}
