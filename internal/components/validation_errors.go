package components

//nolint:stylecheck // ST1001
import (
	"github.com/tksasha/model/errors"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func ValidationErrors(validationErrors errors.ValidationError) Node {
	return If(
		validationErrors.Has("date"),
		Div(
			Ul(
				Class("invalid-feedback"),
				Map(
					validationErrors.Get("date"),
					func(err string) Node {
						return Li(Text(err))
					}),
			),
		),
	)
}
