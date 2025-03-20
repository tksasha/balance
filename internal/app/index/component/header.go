package component

import (
	"github.com/tksasha/balance/internal/common/paths/params"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) header(params params.Params) Node {
	return Header(c.Months(params), c.Years(params))
}
