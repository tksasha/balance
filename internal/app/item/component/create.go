package component

import (
	"github.com/tksasha/balance/internal/app/category"
	"github.com/tksasha/balance/internal/app/item"
	"github.com/tksasha/balance/internal/common/paths/params"
	"github.com/tksasha/validation"
	. "maragu.dev/gomponents" //nolint:staticcheck
	htmx "maragu.dev/gomponents-htmx"
)

func (c *Component) Create(
	params params.Params,
	item *item.Item,
	categories category.Categories,
	errors validation.Errors,
) Node {
	return c.New(
		params,
		categories,
		c.ModalBody(
			c.Form(params, item, categories, errors),
			htmx.SwapOOB("true"),
		),
	)
}
