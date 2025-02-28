package component

import (
	"net/url"

	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) header(values url.Values) Node {
	return Header(c.Months(values), c.Years(values))
}
