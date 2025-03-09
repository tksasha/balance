package component

import (
	. "maragu.dev/gomponents"            //nolint: stylecheck
	. "maragu.dev/gomponents/components" //nolint: stylecheck
	. "maragu.dev/gomponents/html"       //nolint: stylecheck
)

func (c *Component) Input(label, name, value string, classes Classes, message *string) Node {
	if classes == nil {
		classes = make(Classes)
	}

	classes["form-control"] = true
	classes["is-invalid"] = message != nil

	return Div(Class("mb-3"),
		Label(Class("form-label"), Text(label)),
		Input(classes, Name(name), Value(value)),
		Iff(message != nil, func() Node { return Div(Class("invalid-feedback"), Text(*message)) }),
	)
}
