package component

import (
	"github.com/tksasha/balance/internal/backoffice/category"
	"github.com/tksasha/balance/internal/common/paths"
	"github.com/tksasha/balance/internal/common/paths/params"
	"github.com/tksasha/validation"
	. "maragu.dev/gomponents" //nolint: staticcheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint: staticcheck
)

func (c *Component) New(params params.Params, category *category.Category, errors validation.Errors) Node {
	return Div(
		c.Breadcrumbs(
			Li(Class("breadcrumb-item"),
				Span(Class("link"),
					htmx.Get(paths.BackofficeCategories(params)),
					htmx.Target("#modal-body"),
					Text("Категорія"),
				),
			),
			Li(Class("breadcrumb-item active"), Text("Додавання")),
		),
		c.form(category, errors),
	)
}
