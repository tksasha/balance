package component

import (
	"time"

	"github.com/tksasha/balance/internal/common"
)

func (c *Component) date(date time.Time) string {
	if date.IsZero() {
		return ""
	}

	return date.Format(common.DateFormat)
}
