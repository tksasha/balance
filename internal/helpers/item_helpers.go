package helpers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/a-h/templ"
	"github.com/tksasha/balance/internal/models"
)

func ItemsURL(currency *string) templ.SafeURL {
	if currency == nil {
		return templ.URL("/items")
	}

	return templ.URL("/items?currency=" + strings.ToLower(*currency))
}

func ItemURL(item *models.Item) templ.SafeURL {
	return templ.URL(
		"/items/" + strconv.Itoa(item.ID),
	)
}

func EditItemURL(item *models.Item) templ.SafeURL {
	return templ.URL(
		"/items/" + strconv.Itoa(item.ID) + "/edit",
	)
}

func ItemID(item *models.Item) string {
	return fmt.Sprintf("item_%d", item.ID)
}
