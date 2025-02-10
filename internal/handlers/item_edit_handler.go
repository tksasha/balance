package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/responses"
)

type ItemEditHandler struct {
	itemService ItemService
}

func NewItemEditHandler(itemService ItemService) *ItemEditHandler {
	return &ItemEditHandler{
		itemService: itemService,
	}
}

func (h *ItemEditHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.handle(w, r); err != nil {
		if response, ok := w.(*responses.Response); ok {
			response.Error = err

			return
		}

		return
	}
}

func (h *ItemEditHandler) handle(_ http.ResponseWriter, r *http.Request) error {
	item, err := h.itemService.GetItem(r.Context(), r.PathValue("id"))
	if err != nil {
		return err
	}

	_ = item

	return nil
}
