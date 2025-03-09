package component

import (
	"strconv"

	"github.com/tksasha/balance/internal/app/category"
	commoncomponent "github.com/tksasha/balance/internal/common/component"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

type Component struct {
	*commoncomponent.Component
}

func New() *Component {
	return &Component{
		Component: commoncomponent.New(),
	}
}

func (c *Component) itemID(id int) string {
	return "item-" + strconv.Itoa(id)
}

func (c *Component) categories(selected int, categories category.Categories, message *string) Node {
	var nodes []Node

	nodes = append(
		nodes,
		Select(Class("form-select"), Name("category_id"),
			OptGroup(Label(Text("Видатки")),
				Map(categories.Expense(), func(category *category.Category) Node {
					return c.category(category, selected)
				}),
			),
			OptGroup(Label(Text("Надходження")),
				Map(categories.Income(), func(category *category.Category) Node {
					return c.category(category, selected)
				}),
			),
		),
	)
	if message != nil {
		nodes = append(nodes, Div(Class("invalid-feebback"), Text(*message)))
	}

	return c.Map(nodes)
}

func (c *Component) category(category *category.Category, selected int) Node {
	return Option(
		Value(strconv.Itoa(category.ID)),
		Text(category.Name),
		If(category.ID == selected, Selected()),
	)
}
