package component

import (
	"github.com/tksasha/balance/internal/app/item"
	"github.com/tksasha/balance/internal/common/component/path"
	. "maragu.dev/gomponents" //nolint: stylecheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint: stylecheck
)

func (c *Component) Index(items item.Items) Node {
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
		TBody(Map(items, func(item *item.Item) Node { return c.item(item, nil) })),
	)
}

func (c *Component) item(item *item.Item, children ...Node) Node {
	return Tr(ID(c.itemID(item.ID)),
		Td(Class("items-date"),
			Text(c.date(item.Date)),
		),
		Td(Class("items-sum"),
			Div(Class("link"),
				Text(c.Money(item.Sum)),
				htmx.Get(path.EditItem(item.ID)),
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
