package helpers

import (
	"fmt"

	"github.com/tksasha/balance/internal/models"
)

func ItemURL(item models.Item) string {
	return fmt.Sprintf("/items/%d", item.ID)
}
