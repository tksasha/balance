package helpers

import (
	"fmt"

	"github.com/a-h/templ"
	"github.com/tksasha/balance/internal/models"
)

func ItemsURL() templ.SafeURL {
	return templ.URL("/items")
}

func ItemURL(item *models.Item) templ.SafeURL {
	return templ.URL(
		fmt.Sprintf("%s/%d", ItemsURL(), item.ID),
	)
}

func EditItemURL(item *models.Item) templ.SafeURL {
	return templ.URL(
		fmt.Sprintf("%s/edit", ItemURL(item)),
	)
}

func ItemID(item *models.Item) string {
	return fmt.Sprintf("item_%d", item.ID)
}
