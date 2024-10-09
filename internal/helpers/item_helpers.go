package helpers

import (
	"fmt"

	"github.com/a-h/templ"
	"github.com/tksasha/balance/internal/decorators"
	"github.com/tksasha/balance/internal/models"
)

func ItemsURL(currency models.Currency) templ.SafeURL {
	return templ.URL(
		fmt.Sprintf("/%s/items", currency.Name),
	)
}

func ItemURL(currency models.Currency, item *decorators.ItemDecorator) templ.SafeURL {
	return templ.URL(
		fmt.Sprintf("%s/%d", ItemsURL(currency), item.ID),
	)
}

func EditItemURL(currency models.Currency, item *decorators.ItemDecorator) templ.SafeURL {
	return templ.URL(
		fmt.Sprintf("%s/edit", ItemURL(currency, item)),
	)
}

func ItemID(item *decorators.ItemDecorator) string {
	return fmt.Sprintf("item_%d", item.ID)
}
