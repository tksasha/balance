package decorators

import (
	"time"

	"github.com/tksasha/balance/internal/models"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type ItemDecorator struct {
	item    *models.Item
	printer *message.Printer
}

func NewItemDecorator(item *models.Item) *ItemDecorator {
	return &ItemDecorator{
		item: item,
		printer: message.NewPrinter(
			language.Ukrainian,
		),
	}
}

func (d *ItemDecorator) Date() string {
	return d.item.Date.Format(time.DateOnly)
}

func (d *ItemDecorator) Sum() string {
	return d.printer.Sprintf("%.02f", d.item.Sum)
}

func (d *ItemDecorator) Category() string {
	return d.item.Category
}

func (d *ItemDecorator) Description() string {
	return d.item.Description
}
