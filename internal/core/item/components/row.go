package components

import (
	"github.com/tksasha/balance/internal/core/common/components"
	"github.com/tksasha/balance/internal/core/common/helpers"
	"github.com/tksasha/balance/internal/core/item"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func row(helpers *helpers.Helpers, item *item.Item) Node {
	return Tr(
		Td(
			Class("text-center"),
			Text(components.Date(item.Date)),
		),
		Td(
			Class("text-end"),
			A(
				Href(helpers.EditItemPath(item.ID)),
				Text(sum(item.Sum)),
			),
		),
		Td(
			Text(item.CategoryName),
		),
		Td(
			Text(item.Description),
		),
	)
}
