package component

import (
	"github.com/tksasha/balance/internal/backoffice/cash"
	"github.com/tksasha/validation"
	. "maragu.dev/gomponents"      //nolint:staticcheck
	. "maragu.dev/gomponents/html" //nolint:staticcheck
)

func (c *Component) Create(cash *cash.Cash, errors validation.Errors) Node {
	return Div(
		c.form(cash, errors),
	)
}
