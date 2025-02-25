package components

import (
	"github.com/tksasha/balance/internal/backoffice/cash"
	"github.com/tksasha/validator"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *CashComponent) Create(cash *cash.Cash, errors validator.Errors) Node {
	return Div(
		c.form(cash, errors),
	)
}
