package components

import (
	"strconv"

	"github.com/tksasha/balance/internal/core/category"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func option(category *category.Category) Node {
	return Option(
		Value(strconv.Itoa(category.ID)),
		Text(category.Name),
	)
}
