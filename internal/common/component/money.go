package component

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func (c *Component) Money(sum float64) string {
	return message.NewPrinter(language.Ukrainian).Sprintf("%0.2f", sum)
}
