package component

import (
	"github.com/tksasha/balance/internal/backoffice/cash"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) New(cash *cash.Cash) Node {
	return Div(
		c.form(cash, nil),
	)
}
