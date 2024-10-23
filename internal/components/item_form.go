package components

//nolint:stylecheck // ST1001
import (
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/model/errors"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func ItemForm(item *models.Item, validationErrors errors.ValidationError) Node {
	return Form(
		Div(
			Class("mb-3"),
			Label(
				Class("form-label"),
				Text("date"),
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
	)
}
