package helpers

import (
	"fmt"

	"github.com/a-h/templ"
	"github.com/tksasha/balance/internal/decorators"
)

func ItemsURL() templ.SafeURL {
	return templ.URL("/items")
}

func ItemURL(item *decorators.ItemDecorator) templ.SafeURL {
	return templ.URL(
		fmt.Sprintf("%s/%d", ItemsURL(), item.ID),
	)
}

func EditItemURL(item *decorators.ItemDecorator) templ.SafeURL {
	return templ.URL(
		fmt.Sprintf("%s/edit", ItemURL(item)),
	)
}

func ItemID(item *decorators.ItemDecorator) string {
	return fmt.Sprintf("item_%d", item.ID)
}
