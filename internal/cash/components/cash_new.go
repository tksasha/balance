package components

import (
	"github.com/tksasha/balance/internal/cash"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func CashNew(cash *cash.Cash) Node {
	return Div(
		CashForm(cash),
	)
}
