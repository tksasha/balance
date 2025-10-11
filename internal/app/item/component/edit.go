package component

import (
	"github.com/tksasha/balance/internal/app/category"
	"github.com/tksasha/balance/internal/app/item"
	"github.com/tksasha/balance/internal/common/paths/params"
	"github.com/tksasha/validation"
	. "maragu.dev/gomponents" //nolint:staticcheck
)

func (c *Component) Edit(
	params params.Params,
	item *item.Item,
	categories category.Categories,
	errors validation.Errors,
) Node {
	return c.Form(params, item, categories, errors)
}
