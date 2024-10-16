package components

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func IndexPageComponent() g.Node {
	return g.Group{
		g.Raw("<!DOCTYPE html>"),
		h.HTML(
			h.Head(
				h.Title("Balance"),
			),
			h.Body(
				h.H1(g.Text("hello world")),
			),
		),
	}
}
