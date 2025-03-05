package component

import (
	. "maragu.dev/gomponents" //nolint:stylecheck
)

func (c *Component) Map(nodes []Node) Node {
	return Map(nodes, func(node Node) Node { return node })
}
