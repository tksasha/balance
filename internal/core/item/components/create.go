package components

import (
	"github.com/tksasha/balance/internal/core/item"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func Create(item *item.Item) Node {
	return Div(
		Text(item.Description),
	)
}
