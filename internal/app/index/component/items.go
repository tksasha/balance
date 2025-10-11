package component

import (
	. "maragu.dev/gomponents"      //nolint:staticcheck
	. "maragu.dev/gomponents/html" //nolint:staticcheck
)

func (c *Component) items() Node {
	return Div(Class("container"),
		Div(Class("row mt-4 mb-5"),
			Div(Class("col"),
				Div(Class("card items"),
					Div(Class("card-body"), ID("items")),
				),
			),
		),
	)
}
