package requests

import (
	"github.com/tksasha/balance/pkg/requestvalidator"
)

type CreateItemRequest struct {
	*requestvalidator.Base

	Date string
}

func NewCreateItemRequest(date string) *CreateItemRequest {
	req := new(CreateItemRequest)

	req.Base = requestvalidator.New()

	return req
}

func (r *CreateItemRequest) Validate() {
	if r.Date == "" {
		r.AddError("date", "is required")
	}

	r.AddError("date", "cannot be blank")
}
