package component

import (
	. "maragu.dev/gomponents" //nolint:stylecheck
)

func (c *Component) Template(children ...Node) Node {
	return El("template", children...)
}
