package components

//nolint:stylecheck // ST1001
import (
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/model/errors"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

//nolint:funlen
func ItemForm(item *models.Item, categories models.Categories, validationErrors errors.ValidationError) Node {
	return Form(
		Div(
			Class("mb-3"),
			Label(
				Class("form-label"),
				Text("Date"),
			),
			Input(
				If(
					validationErrors.Has("date"),
					Class("form-control invalid"),
				),
				If(
					!validationErrors.Has("date"),
					Class("form-control"),
				),
				Value(item.GetDateAsString()),
			),
			ValidationErrors(validationErrors),
		),
		Div(
			Class("mb-3"),
			Label(
				Class("form-label"),
				Text("Sum"),
			),
			Input(
				If(
					validationErrors.Has("formula"),
					Class("form-control invalid"),
				),
				If(
					!validationErrors.Has("formula"),
					Class("form-control"),
				),
				Value(item.GetSumAsString()),
			),
		),
		Div(
			Class("mb-3"),
			Label(
				Class("form-label"),
				Text("Category"),
			),
			Select(
				Class("form-select"),
				OptGroup(
					Attr("Label", "Expense"),
					Map(categories.Expense(), func(category *models.Category) Node {
						return Option(
							Value(category.GetIDAsString()),
							Text(category.Name),
						)
					}),
				),
				OptGroup(
					Attr("Label", "Income"),
					Map(categories.Income(), func(category *models.Category) Node {
						return Option(
							Value(category.GetIDAsString()),
							Text(category.Name),
						)
					}),
				),
			),
		),
	)
}
