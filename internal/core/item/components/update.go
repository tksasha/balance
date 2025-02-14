package components

import (
	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/core/item"
	"github.com/tksasha/balance/pkg/validation"
	. "maragu.dev/gomponents" //nolint:stylecheck
)

func Update(item *item.Item, categories category.Categories, errors validation.Errors) Node {
	return form(item, categories, errors)
}
