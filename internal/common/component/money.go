package component

import (
	"github.com/shopspring/decimal"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func (c *Component) Money(sum decimal.Decimal) string {
	f, _ := sum.Float64()

	return message.NewPrinter(language.Ukrainian).Sprintf("%0.2f", f)
}
