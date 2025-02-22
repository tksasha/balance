package component

import (
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) Modal() Node {
	return Div(
		ID("modal"),
		Class("modal modal-blur fade"),
		Style("display: none"),
		TabIndex("-1"),
		c.ModalDialog(),
	)
}

func (c *Component) ModalDialog() Node {
	return Div(
		Class("modal-dialog modal-lg modal-dialog-centered"),
		Div(
			Class("modal-content"),
			Div(
				ID("modal-body"),
				Class("modal-body"),
			),
		),
	)
}
