package component

import (
	"github.com/tksasha/balance/internal/app/item"
	"github.com/tksasha/balance/internal/common/paths/params"
	. "maragu.dev/gomponents" //nolint:staticcheck
	htmx "maragu.dev/gomponents-htmx"
)

func (c *Component) Update(params params.Params, item *item.Item) Node {
	return c.Template(c.item(params, item, htmx.SwapOOB("true")))
}
