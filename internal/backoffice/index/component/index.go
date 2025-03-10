package component

import (
	"github.com/tksasha/balance/internal/common/component/path"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) Index() Node {
	return c.BackofficeModalBody(
		Div(
			Div(Class("d-inline-flex me-3"),
				Button(Class("btn btn-outline-primary btn-lg"),
					htmx.Get(path.BackofficeCashes()),
					htmx.Target("#backoffice-modal"),
					Text("Залишки"),
				),
			),
			Div(Class("d-inline-flex"),
				Button(Class("btn btn-outline-primary btn-lg"),
					htmx.Get(path.BackofficeCategories()),
					htmx.Target("#backoffice-modal"),
					Text("Категорії"),
				),
			),
		),
	)
}
