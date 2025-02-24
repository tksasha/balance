package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/common"
	"github.com/tksasha/balance/internal/core/item"
)

type DeleteHandler struct {
	*common.BaseHandler

	itemService item.Service
}

func NewDeleteHandler(
	baseHandler *common.BaseHandler,
	itemService item.Service,
) *DeleteHandler {
	return &DeleteHandler{
		BaseHandler: baseHandler,
		itemService: itemService,
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
	return h.itemService.Delete(r.Context(), r.PathValue("id"))
}
