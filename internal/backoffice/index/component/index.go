package component

import (
	"github.com/tksasha/balance/internal/common/paths"
	"github.com/tksasha/balance/internal/common/paths/params"
	. "maragu.dev/gomponents" //nolint:staticcheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint:staticcheck
)

func (c *Component) Index(params params.Params) Node {
	return c.ModalBody(
		Div(Class("row"),
			Div(Class("col mb-3"),
				Button(Class("btn btn-outline-primary btn-lg w-100"),
					htmx.Get(paths.BackofficeCashes(params)),
					htmx.Target("#modal-body"),
					Text("Залишки"),
				),
			),
			Div(Class("col"),
				Button(Class("btn btn-outline-primary btn-lg w-100"),
					htmx.Get(paths.BackofficeCategories(params)),
					htmx.Target("#modal-body"),
					Text("Категорії"),
				),
			),
		),
	)
}
