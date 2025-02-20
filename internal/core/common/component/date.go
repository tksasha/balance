package component

import "time"

func (c *Component) Date(date time.Time) string {
	if date.IsZero() {
		return ""
	}

	return date.Format(time.DateOnly)
}
