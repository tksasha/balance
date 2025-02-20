package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/core/cash"
	"github.com/tksasha/balance/internal/core/common"
)

type DeleteHandler struct {
	*common.BaseHandler

	cashService cash.Service
}

func NewDeleteHandler(
	baseHandler *common.BaseHandler,
	cashService cash.Service,
) *DeleteHandler {
	return &DeleteHandler{
		BaseHandler: baseHandler,
		cashService: cashService,
	}
}

func (h *DeleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.handle(r); err != nil {
		h.SetError(w, err)

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *DeleteHandler) handle(r *http.Request) error {
	return h.cashService.Delete(r.Context(), r.PathValue("id"))
}
