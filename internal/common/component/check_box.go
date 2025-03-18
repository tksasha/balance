package component

import (
	. "maragu.dev/gomponents"      //nolint: stylecheck
	. "maragu.dev/gomponents/html" //nolint: stylecheck
)

func (c *Component) CheckBox(label, id string, value bool) Node {
	return Div(Class("mb-3"),
		Div(Class("form-check"),
			Input(Class("form-check-input"), Type("checkbox"), ID(id),
				If(value, Checked()),
			),
			Label(Class("form-check-label"), For(id), Text(label)),
		),
	)
}
