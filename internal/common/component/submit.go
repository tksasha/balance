package component

import (
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) Submit(id int) Node {
	return Div(Class("mb-3"),
		Button(Class("btn btn-primary"),
			If(id == 0, Text("Створити")),
			If(id != 0, Text("Оновити")),
		),
	)
}
