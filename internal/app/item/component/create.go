package component

import (
	"github.com/tksasha/balance/internal/app/category"
	"github.com/tksasha/balance/internal/app/item"
	"github.com/tksasha/validation"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) Create(item *item.Item, categories category.Categories, errors validation.Errors) Node {
	_ = item
	_ = categories
	_ = errors

	return Form()
}
