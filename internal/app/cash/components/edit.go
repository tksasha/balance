package components

import (
	"github.com/tksasha/balance/internal/app/cash"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *CashComponent) Edit(cash *cash.Cash) Node {
	return Div(
		c.form(cash, nil),
	)
}
