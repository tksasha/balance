package handler

import (
	"net/http"

	"github.com/tksasha/balance/internal/common"
)

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) SetError(w http.ResponseWriter, err error) {
	if wrapper, ok := w.(*common.ResponseWriterWrapper); ok {
		wrapper.Error = err
	}
}
