package component

import (
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) No() Node {
	return Div(Class("text-secondary"), Text("ні"))
}
