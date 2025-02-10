package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/apperrors"
	"github.com/tksasha/balance/internal/requests"
	"github.com/tksasha/balance/internal/responses"
)

type ItemUpdateHandler struct {
	itemService ItemService
}

func NewItemUpdateHandler(itemService ItemService) *ItemUpdateHandler {
	return &ItemUpdateHandler{
		itemService: itemService,
	}
}

func (h *ItemUpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.handle(r); err != nil {
		if response, ok := w.(*responses.Response); ok {
			response.Error = err

			return
		}

		return
	}

	_, _ = w.Write([]byte("render update page\n"))
}

func (h *ItemUpdateHandler) handle(r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return apperrors.ErrParsingForm
	}

	return h.itemService.Update(
		r.Context(),
		requests.ItemUpdateRequest{
			ID:          r.PathValue("id"),
			Date:        r.FormValue("date"),
			Formula:     r.FormValue("formula"),
			CategoryID:  r.FormValue("category_id"),
			Description: r.FormValue("description"),
		},
	)
}
