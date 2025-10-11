package component

import (
	. "maragu.dev/gomponents"      //nolint:staticcheck
	. "maragu.dev/gomponents/html" //nolint:staticcheck
)

func (c *Component) categories() Node {
	return Div(Class("container-fluid"), ID("categories"))
}
