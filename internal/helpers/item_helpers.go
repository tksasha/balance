package helpers

import (
	"fmt"

	"github.com/a-h/templ"
	"github.com/tksasha/balance/internal/models"
)

func EditItemURL(item *models.Item) templ.SafeURL {
	return templ.URL(
		fmt.Sprintf("/items/%d/edit", item.ID),
	)
}
