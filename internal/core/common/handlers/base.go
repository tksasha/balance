package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/core/common/response"
)

type BaseHandler struct{}

func NewBaseHandler() *BaseHandler {
	return &BaseHandler{}
}

func (h *BaseHandler) SetError(w http.ResponseWriter, err error) {
	if wrapper, ok := w.(*response.Wrapper); ok {
		wrapper.Error = err
	}
}
