package handler

import (
	"net/http"

	"github.com/tksasha/balance/internal/common/middleware"
)

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) SetError(w http.ResponseWriter, err error) {
	if wrapper, ok := w.(*middleware.ResponseWriterWrapper); ok {
		wrapper.Error = err
	}
}
