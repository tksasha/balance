package component

import (
	. "maragu.dev/gomponents"      //nolint: staticcheck
	. "maragu.dev/gomponents/html" //nolint: staticcheck
)

func (c *Component) CheckBox(label, name, id string, value bool) Node {
	return Div(Class("mb-3"),
		Div(Class("form-check"),
			Input(Class("form-check-input"), Type("checkbox"), Name(name), ID(id), Value("true"),
				If(value, Checked()),
			),
			Label(Class("form-check-label"), For(id), Text(label)),
		),
	)
}
