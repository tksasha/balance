package requests

import (
	"time"

	"github.com/tksasha/formula"
	"github.com/tksasha/model"
	"github.com/tksasha/model/errors/messages"
)

type CreateItemRequest struct {
	*model.Model

	Date        time.Time
	Formula     string
	Sum         float64
	Description string
}

func NewCreateItemRequest() *CreateItemRequest {
	req := new(CreateItemRequest)

	req.Model = model.New()

	return req
}

func (r *CreateItemRequest) GetDate() string {
	if r.Date.IsZero() {
		return ""
	}

	return r.Date.Format(time.DateOnly)
}

func (r *CreateItemRequest) Parse(
	inputDate,
	inputFormula,
	inputCategoryID,
	inputDescription string,
) {
	var date time.Time

	var err error

	date, err = time.Parse(time.DateOnly, inputDate)
	if err != nil {
		r.Errors.Set("date", messages.Invalid)
	}

	if date.IsZero() {
		r.Errors.Set("date", messages.Required)
	}

	if inputFormula == "" {
		r.Errors.Set("formula", messages.Required)
	}

	r.Date = date
	r.Formula = inputFormula

	r.Sum, err = formula.Calculate(inputFormula)
	if err != nil {
		r.Errors.Set("formula", messages.Invalid)
	}

	_ = inputCategoryID

	r.Description = inputDescription
}
