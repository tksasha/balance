package requests

import (
	"strconv"
	"time"

	calculator "github.com/tksasha/formula"
	"github.com/tksasha/model"
	"github.com/tksasha/model/errors/messages"
)

type CreateItemRequest struct {
	*model.Model

	Date        time.Time
	Formula     string
	Sum         float64
	CategoryID  int
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
	date,
	formula,
	categoryID,
	description string,
) {
	var err error

	r.Date, err = time.Parse(time.DateOnly, date)
	if err != nil {
		r.Errors.Set("date", messages.Invalid)
	}

	if r.Date.IsZero() {
		r.Errors.Set("date", messages.Required)
	}

	r.Formula = formula

	if r.Formula == "" {
		r.Errors.Set("formula", messages.Required)
	}

	r.Sum, err = calculator.Calculate(r.Formula)
	if err != nil {
		r.Errors.Set("formula", messages.Invalid)
	}

	if categoryID == "" {
		r.Errors.Set("category-id", messages.Required)
	}

	r.CategoryID, err = strconv.Atoi(categoryID)
	if err != nil {
		r.Errors.Set("category-id", messages.Invalid)
	}

	if r.CategoryID == 0 {
		r.Errors.Set("category-id", messages.Zero)
	}

	r.Description = description
}
