package component

import (
	"github.com/tksasha/balance/internal/app/item"
	"github.com/tksasha/balance/internal/common/paths"
	"github.com/tksasha/balance/internal/common/paths/params"
	. "maragu.dev/gomponents" //nolint: stylecheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint: stylecheck
)

func (c *Component) Index(params params.Params, items item.Items) Node {
	if len(items) == 0 {
		return Div(Class("text-success text-center p-2"), Text("Нічого не знайдено"))
	}

	return Table(
		Class("table table-hover table-borderless"),
		THead(
			Tr(
				Th(Class("items-date"), Text("Дата")),
				Th(Class("items-sum"), Text("Сума")),
				Th(Class("items-category"), Text("Категорія")),
				Th(Text("Опис")),
			),
		),
		TBody(Map(items, func(item *item.Item) Node { return c.item(params, item, nil) })),
	)
}

func (c *Component) item(params params.Params, item *item.Item, children ...Node) Node {
	return Tr(ID(c.itemID(item.ID)),
		Td(Class("items-date"),
			Text(c.date(item.Date)),
		),
		Td(Class("items-sum"),
			Div(Class("link"),
				Text(c.Money(item.Sum)),
				htmx.Get(paths.EditItem(params, item.ID)),
				htmx.Target("#modal-body"),
				htmx.Trigger("click"),
				Data("bs-toggle", "modal"),
				Data("bs-target", "#modal"),
			),
		),
		Td(Class("items-category"),
			Text(item.CategoryName),
		),
		Td(c.Description(item.Description)),
		Group(children),
	)
}
