package component

import (
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
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

func (c *Component) BackofficeModal() Node {
	return Div(ID("backoffice-modal"), Class("modal modal-blur fade"),
		Style("display: none"),
		TabIndex("-1"),
		Div(Class("modal-dialog modal-xl modal-dialog-centered"),
			Div(Class("modal-content"),
				c.BackofficeModalBody(),
			),
		),
	)
}

func (c *Component) BackofficeModalBody(children ...Node) Node {
	return Div(ID("backoffice-modal-body"), Class("modal-body"),
		Group(children),
	)
}
