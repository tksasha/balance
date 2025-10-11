package component

import (
	"strconv"

	. "maragu.dev/gomponents" //nolint:staticcheck
)

func (c *Component) Text(el any) Node {
	if i, ok := el.(int); ok {
		return Text(strconv.Itoa(i))
	}

	if s, ok := el.(string); ok {
		return Text(s)
	}

	return Text("NONE")
}
