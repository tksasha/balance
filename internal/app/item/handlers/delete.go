package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/app/item"
	"github.com/tksasha/balance/internal/common/handler"
)

type DeleteHandler struct {
	*handler.Handler

	itemService item.Service
}

func NewDeleteHandler(
	itemService item.Service,
) *DeleteHandler {
	return &DeleteHandler{
		Handler:     handler.New(),
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
