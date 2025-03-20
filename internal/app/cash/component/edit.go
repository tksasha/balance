package component

import (
	"github.com/tksasha/balance/internal/app/cash"
	"github.com/tksasha/balance/internal/common/paths"
	"github.com/tksasha/balance/internal/common/paths/params"
	"github.com/tksasha/validation"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
)

func (c *Component) Edit(params params.Params, cash *cash.Cash, errors validation.Errors) Node {
	return c.form(
		Group([]Node{
			htmx.Patch(paths.UpdateCash(params, cash.ID)),
			htmx.Target("#modal-body"),
			htmx.Swap("outerHTML"),
		}),
		cash,
		errors,
	)
}
