package components

import (
	"fmt"

	"github.com/tksasha/balance/internal/core/common"
	"github.com/tksasha/balance/internal/core/common/components"
	"github.com/tksasha/balance/internal/core/item"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func editLink(item *item.Item) Node {
	return A(
		Href(
			fmt.Sprintf("%s/%d/edit", common.Items, item.ID),
		),
		Text(
			components.Sum(item.Sum),
		),
	)
}
