package components

import (
	"github.com/tksasha/model/errors"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func ValidationErrors(validationErrors errors.ValidationError) gomponents.Node {
	return gomponents.If(
		validationErrors.Has("date"),
		html.Ul(
			html.Class("invalid-feedback"),
			gomponents.Map(
				validationErrors.Get("date"),
				func(err string) gomponents.Node {
					return html.Li(gomponents.Text(err))
				}),
		),
	)
}
