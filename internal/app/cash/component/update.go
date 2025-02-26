package component

import (
	"github.com/tksasha/balance/internal/app/cash"
	. "maragu.dev/gomponents" //nolint:stylecheck
)

func (c *CashComponent) Update(cash *cash.Cash) Node {
	return c.Cash(cash)
}
