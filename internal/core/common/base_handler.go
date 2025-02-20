package common

import (
	"net/http"
)

type BaseHandler struct{}

func NewBaseHandler() *BaseHandler {
	return &BaseHandler{}
}

func (h *BaseHandler) SetError(w http.ResponseWriter, err error) {
	if wrapper, ok := w.(*ResponseWriterWrapper); ok {
		wrapper.Error = err
	}
}
