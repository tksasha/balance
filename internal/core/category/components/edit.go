package components

import (
	"github.com/tksasha/balance/internal/core/category"
	. "maragu.dev/gomponents"      //nolint: stylecheck
	. "maragu.dev/gomponents/html" //nolint: stylecheck
)

func Edit(category *category.Category) Node {
	return Div(
		form(category),
	)
}
