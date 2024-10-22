package components

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Container(child gomponents.Node) gomponents.Node {
	return html.Div(html.Class("container"), child)
}
