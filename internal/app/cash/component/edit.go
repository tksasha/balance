package component

import (
	"net/url"

	"github.com/tksasha/balance/internal/app/cash"
	"github.com/tksasha/balance/internal/common/component/path"
	"github.com/tksasha/validation"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
)

func (c *Component) Edit(values url.Values, cash *cash.Cash, errors validation.Errors) Node {
	return c.form(
		Group([]Node{
			htmx.Patch(path.UpdateCash(values, cash.ID)),
			htmx.Target("#modal-body"),
			htmx.Swap("outerHTML"),
		}),
		cash,
		errors,
	)
}
