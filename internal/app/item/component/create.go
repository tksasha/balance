package component

import (
	"net/url"

	"github.com/tksasha/balance/internal/app/category"
	"github.com/tksasha/balance/internal/app/item"
	"github.com/tksasha/validation"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
)

func (c *Component) Create(
	values url.Values,
	item *item.Item,
	categories category.Categories,
	errors validation.Errors,
) Node {
	return c.New(
		values,
		categories,
		c.ModalBody(
			c.Form(values, item, categories, errors),
			htmx.SwapOOB("true"),
		),
	)
}
