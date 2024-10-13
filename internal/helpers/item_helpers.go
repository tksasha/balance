package helpers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/a-h/templ"
	"github.com/tksasha/balance/internal/decorators"
	"github.com/tksasha/balance/internal/models"
)

func ItemsURL(currency *models.Currency) templ.SafeURL {
	if currency == nil {
		return templ.URL("/items")
	}

	return templ.URL("/items?currency=" + strings.ToLower(currency.Code))
}

func ItemURL(item *decorators.ItemDecorator) templ.SafeURL {
	return templ.URL(
		"/items/" + strconv.Itoa(item.ID),
	)
}

func EditItemURL(item *decorators.ItemDecorator) templ.SafeURL {
	return templ.URL(
		"/items/" + strconv.Itoa(item.ID) + "/edit",
	)
}

func ItemID(item *decorators.ItemDecorator) string {
	return fmt.Sprintf("item_%d", item.ID)
}
