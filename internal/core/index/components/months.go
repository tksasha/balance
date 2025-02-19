package components

import (
	"fmt"
	"net/http"

	"github.com/tksasha/balance/internal/core/common/helpers"
	"github.com/tksasha/month"
	. "maragu.dev/gomponents" //nolint:stylecheck
	hx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func Months(helpers *helpers.Helpers, request *http.Request) Node {
	return Div(
		Map(month.All(), func(month month.Month) Node {
			return A(
				If(
					isActive(request.URL.Query().Get("month"), month.Number),
					Class("active"),
				),
				Href(helpers.ItemsPath(request, 0, month.Number)),
				Text(month.Name),
				hx.Get(helpers.ItemsPath(request, 0, month.Number)),
			)
		}),
	)
}

func Month(helpers *helpers.Helpers, request *http.Request) Node {
	return Div()
}

func isActive(active string, number int) bool {
	return active == fmt.Sprintf("%02d", number)
}
