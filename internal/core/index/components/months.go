package components

import (
	"fmt"
	"net/http"

	"github.com/tksasha/balance/internal/core/common/helpers"
	. "maragu.dev/gomponents" //nolint:stylecheck
	hx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

type month struct {
	number int
	name   string
}

var months = []month{ //nolint:gochecknoglobals
	{1, "Січень"},
	{2, "Лютий"},
	{3, "Березень"},
	{4, "Квітень"},
	{5, "Травень"},
	{6, "Червень"},
	{7, "Липень"},
	{8, "Серпень"},
	{9, "Вересень"},
	{10, "Жовтень"},
	{11, "Листопад"},
	{12, "Грудень"},
}

func Months(helpers *helpers.Helpers, request *http.Request) Node {
	return Div(
		Map(months, func(month month) Node {
			return A(
				If(
					isActive(request.URL.Query().Get("month"), month.number),
					Class("active"),
				),
				Href(helpers.ItemsPath(request, 0, month.number)),
				Text(month.name),
				hx.Get(helpers.ItemsPath(request, 0, month.number)),
			)
		}),
	)
}

func isActive(active string, number int) bool {
	return active == fmt.Sprintf("%02d", number)
}
