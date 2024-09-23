package decorators

import (
	"time"

	"github.com/tksasha/balance/internal/models"
)

type ItemDecorator struct {
	item *models.Item
}

func NewItemDecorator(item *models.Item) *ItemDecorator {
	return &ItemDecorator{
		item: item,
	}
}

func (d *ItemDecorator) Date() string {
	return d.item.Date.Format(time.DateOnly)
}

func (d *ItemDecorator) Sum() float32 {
	return d.item.Sum
}

func (d *ItemDecorator) Category() string {
	return d.item.Category
}

func (d *ItemDecorator) Description() string {
	return d.item.Description
}
