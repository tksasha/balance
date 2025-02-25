package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/backoffice/category"
	"github.com/tksasha/balance/internal/common/handler"
)

type DeleteHandler struct {
	*handler.Handler

	categoryService category.Service
}

func NewDeleteHandler(
	categoryService category.Service,
) *DeleteHandler {
	return &DeleteHandler{
		Handler:         handler.New(),
		categoryService: categoryService,
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
	return h.categoryService.Delete(r.Context(), r.PathValue("id"))
}
