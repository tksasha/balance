package components

import (
	"github.com/tksasha/balance/internal/item"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func Update(item *item.Item) Node {
	return Div(
		Text(item.Description),
	)
}
