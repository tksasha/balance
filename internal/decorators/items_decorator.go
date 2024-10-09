package decorators

import "github.com/tksasha/balance/internal/models"

type ItemsDecorator struct {
	Items []*ItemDecorator
}

func NewItemsDecorator(items []*models.Item) *ItemsDecorator {
	decorator := &ItemsDecorator{}

	for _, item := range items {
		decorator.Items = append(decorator.Items, NewItemDecorator(item))
	}

	return decorator
}
