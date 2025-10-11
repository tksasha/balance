package component

import (
	. "maragu.dev/gomponents"      //nolint:staticcheck
	. "maragu.dev/gomponents/html" //nolint:staticcheck
)

func (c *Component) Modal() Node {
	return Div(ID("modal"), Class("modal modal-blur fade"),
		Style("display: none"),
		TabIndex("-1"),
		Div(Class("modal-dialog modal-lg modal-dialog-centered"),
			Div(Class("modal-content"),
				c.ModalBody(nil),
			),
		),
	)
}

func (c *Component) ModalBody(children ...Node) Node {
	return Div(ID("modal-body"), Class("modal-body"),
		Group(children),
	)
}
