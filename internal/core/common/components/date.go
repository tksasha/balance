package components

import "time"

func (c *BaseComponent) Date(date time.Time) string {
	if date.IsZero() {
		return ""
	}

	return date.Format(time.DateOnly)
}
