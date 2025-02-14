package components

import (
	"github.com/tksasha/balance/internal/core/cash"
	"github.com/tksasha/balance/pkg/validation"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func Create(cash *cash.Cash, errors validation.Errors) Node {
	return Div(
		form(cash, errors),
	)
}
