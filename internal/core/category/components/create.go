package components

import (
	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/pkg/validation"
	. "maragu.dev/gomponents"      //nolint: stylecheck
	. "maragu.dev/gomponents/html" //nolint: stylecheck
)

func Create(category *category.Category, errors validation.Errors) Node {
	return Div(
		form(category, errors),
	)
}
