package components

import (
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/model/errors"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func ItemForm(item *models.Item, validationErrors errors.ValidationError) gomponents.Node {
	return html.Form(
		html.Div(
			html.Class("mb-3"),
			html.Label(
				html.Class("form-label"),
				gomponents.Text("date"),
			),
			html.Input(
				gomponents.If(
					validationErrors.Has("date"),
					html.Class("form-control invalid"),
				),
				gomponents.If(
					!validationErrors.Has("date"),
					html.Class("form-control"),
				),
				html.Value(item.GetDateAsString()),
			),
			html.Div(ValidationErrors(validationErrors)),
		),
	)
}
