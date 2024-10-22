package components

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/components"
)

func IndexPage() gomponents.Node {
	return components.HTML5(
		components.HTML5Props{
			Title:    "Balance",
			Language: "en",
			Head: []gomponents.Node{
				Stylesheet("/assets/bootstrap.min.css"),
				Stylesheet("/assets/application.css"),
			},
			Body: []gomponents.Node{
				Container(
					Card(
						gomponents.Text("we are here"),
					),
				),
			},
		},
	)
}
