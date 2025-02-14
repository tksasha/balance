package components

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Stylesheet(path string) gomponents.Node {
	return html.Link(html.Rel("stylesheet"), html.Href(path))
}
