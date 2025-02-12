package components

import (
	"github.com/tksasha/balance/internal/models"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func CashNew(cash *models.Cash) Node {
	return Div(
		Form(),
	)
}
