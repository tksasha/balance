package components

import (
	"github.com/tksasha/balance/internal/category"
	. "maragu.dev/gomponents"      //nolint: stylecheck
	. "maragu.dev/gomponents/html" //nolint: stylecheck
)

func Update(category *category.Category) Node {
	return Div(
		Text(category.Name),
	)
}
