package component

import (
	"github.com/tksasha/balance/internal/common/component/path"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) Index() Node {
	return c.ModalBody(
		Div(Class("row"),
			Div(Class("col mb-3"),
				Button(Class("btn btn-outline-primary btn-lg w-100"),
					htmx.Get(path.BackofficeCashes(nil)),
					htmx.Target("#modal-body"),
					Text("Залишки"),
				),
			),
			Div(Class("col"),
				Button(Class("btn btn-outline-primary btn-lg w-100"),
					htmx.Get(path.BackofficeCategories()),
					htmx.Target("#backoffice-modal-body"),
					Text("Категорії"),
				),
			),
		),
	)
}
