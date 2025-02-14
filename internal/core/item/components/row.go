package components

import (
	"time"

	"github.com/tksasha/balance/internal/core/item"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func row(item *item.Item) Node {
	return Tr(
		Td(Class("text-center"), Text(item.Date.Format(time.DateOnly))),
		Td(
			Class("text-end"),
			A(
				Href("/items"),
				Text("sum should be here"),
			),
		),
		Td(Text(item.CategoryName)),
		Td(Text(item.Description)),
	)
}
