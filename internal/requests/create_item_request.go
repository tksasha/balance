package requests

import (
	"time"

	"github.com/tksasha/model"
)

type CreateItemRequest struct {
	*model.Model

	Formula string
	date    time.Time
}

func NewCreateItemRequest() *CreateItemRequest {
	req := new(CreateItemRequest)

	req.Model = model.New()

	return req
}

func (r *CreateItemRequest) Date() string {
	if r.date.IsZero() {
		return ""
	}

	return r.date.Format(time.DateOnly)
}

func (r *CreateItemRequest) SetDate(date time.Time) {
	r.date = date
}

func (r *CreateItemRequest) Parse(
	inputDate,
	formula string,
) {
	date, err := time.Parse(time.DateOnly, inputDate)
	if err != nil {
		r.Errors.Set("date", "is invalid")
	}

	if date.IsZero() {
		r.Errors.Set("date", "is required")
	}

	if formula == "" {
		r.Errors.Set("formula", "is required")
	}

	r.SetDate(date)
	r.Formula = formula
}
