package components

import (
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/model/errors"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func ValidationErrors(validationErrors errors.ValidationError) g.Node {
	return g.If(
		validationErrors.Has("date"),
		h.Ul(
			h.Class("invalid-feedback"),
			g.Map(
				validationErrors.Get("date"),
				func(err string) g.Node {
					return h.Li(g.Text(err))
				}),
		),
	)
}

func ItemForm(item *models.Item, validationErrors errors.ValidationError) g.Node {
	return h.Form(
		h.Div(
			h.Class("mb-3"),
			h.Label(
				h.Class("form-label"),
				g.Text("date"),
			),
			h.Input(
				g.If(
					validationErrors.Has("date"),
					h.Class("form-control invalid"),
				),
				g.If(
					!validationErrors.Has("date"),
					h.Class("form-control"),
				),
				h.Value(item.GetDateAsString()),
			),
			h.Div(ValidationErrors(validationErrors)),
		),
	)
}
