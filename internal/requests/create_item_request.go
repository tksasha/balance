package requests

import (
	"github.com/tksasha/model"
)

type CreateItemRequest struct {
	*model.Model

	Date string
}

func NewCreateItemRequest(date string) *CreateItemRequest {
	req := new(CreateItemRequest)

	req.Model = model.New()

	return req
}

func (r *CreateItemRequest) Validate() {
	if r.Date == "" {
		r.Errors.Set("date", "is required")
	}

	r.Errors.Set("date", "cannot be blank")
}
