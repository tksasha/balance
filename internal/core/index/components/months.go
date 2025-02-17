package components

import (
	"net/http"

	"github.com/tksasha/balance/internal/core/common/helpers"
	"github.com/tksasha/balance/internal/core/index"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func Months(index *index.Index, request *http.Request) Node {
	return Div(
		Map(index.Months, func(name string) Node {
			return month(name, request)
		}),
	)
}

func month(name string, request *http.Request) Node {
	return A(
		Href(helpers.ItemsPath(request, 0, 0)),
		Text(name),
	)
}
