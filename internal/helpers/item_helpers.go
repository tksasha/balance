package helpers

import (
	"fmt"

	"github.com/tksasha/balance/internal/models"
)

func EditItemURL(item models.Item) string {
	return fmt.Sprintf("/items/%d/edit", item.ID)
}
