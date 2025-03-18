package component

import (
	"github.com/tksasha/balance/internal/backoffice/category"
	"github.com/tksasha/balance/internal/common/component/path"
	"github.com/tksasha/validation"
	. "maragu.dev/gomponents" //nolint: stylecheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint: stylecheck
)

func (c *Component) New(params path.Params, category *category.Category, errors validation.Errors) Node {
	return Div(
		c.Breadcrumbs(
			Li(Class("breadcrumb-item"),
				Span(Class("link"),
					htmx.Get(path.BackofficeCategories(params)),
					htmx.Target("#modal-body"),
					Text("Категорія"),
				),
			),
			Li(Class("breadcrumb-item active"), Text("Додавання")),
		),
		c.form(category, errors),
	)
}
