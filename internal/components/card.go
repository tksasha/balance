package components

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Card(child gomponents.Node) gomponents.Node {
	return html.Div(
		html.Class("card"),
		html.Div(
			html.Class("card-body"),
			child,
		),
	)
}
