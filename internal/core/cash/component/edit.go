package component

import (
	"github.com/tksasha/balance/internal/core/cash"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) Edit(cash *cash.Cash) Node {
	return Div(
		c.form(cash, nil),
	)
}
