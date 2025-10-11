package component

import (
	. "maragu.dev/gomponents" //nolint:staticcheck
)

func (c *Component) Template(children ...Node) Node {
	return El("template", children...)
}
