package handlers

import (
	"net/http"

	apperrors "github.com/tksasha/balance/internal/core/common"
	"github.com/tksasha/balance/internal/core/common/handlers"
	"github.com/tksasha/balance/internal/core/item"
	"github.com/tksasha/balance/internal/core/item/components"
)

type UpdateHandler struct {
	service item.Service
}

func NewUpdateHandler(service item.Service) *UpdateHandler {
	return &UpdateHandler{
		service: service,
	}
}

func (h *UpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	item, err := h.handle(r)
	if err != nil {
		handlers.E(w, err)

		return
	}

	if err := components.Update(item).Render(w); err != nil {
		handlers.E(w, err)
	}
}

func (h *UpdateHandler) handle(r *http.Request) (*item.Item, error) {
	if err := r.ParseForm(); err != nil {
		return nil, apperrors.ErrParsingForm
	}

	request := item.UpdateRequest{
		ID:          r.PathValue("id"),
		Date:        r.FormValue("date"),
		Formula:     r.FormValue("formula"),
		CategoryID:  r.FormValue("category_id"),
		Description: r.FormValue("description"),
	}

	return h.service.Update(r.Context(), request)
}
