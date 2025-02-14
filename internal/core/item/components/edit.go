package components

import (
	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/core/item"
	. "maragu.dev/gomponents" //nolint:stylecheck
)

func Edit(item *item.Item, categories category.Categories) Node {
	return form(item, categories, nil)
}
