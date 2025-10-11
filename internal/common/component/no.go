package component

import (
	. "maragu.dev/gomponents"      //nolint:staticcheck
	. "maragu.dev/gomponents/html" //nolint:staticcheck
)

func (c *Component) No() Node {
	return Div(Class("text-secondary"), Text("ні"))
}
