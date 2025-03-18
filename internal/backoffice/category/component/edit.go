package component

import (
	"github.com/tksasha/balance/internal/backoffice/category"
	"github.com/tksasha/balance/internal/common/component/path"
	. "maragu.dev/gomponents" //nolint: stylecheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint: stylecheck
)

func (c *Component) Edit(params path.Params, category *category.Category) Node {
	return Div(
		c.Breadcrumbs(
			Li(Class("breadcrumb-item"),
				Span(Class("link"),
					htmx.Get(path.BackofficeCategories(params)),
					htmx.Target("#modal-body"),
					Text("Категорії"),
				),
			),
			Li(Class("breadcrumb-item active"), Text("Редагування")),
		),
		c.form(category, nil),
	)
}
