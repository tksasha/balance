package handlers

import (
	"fmt"
	"net/http"

	"github.com/tksasha/balance/internal/app/item"
	"github.com/tksasha/balance/internal/app/item/component"
	"github.com/tksasha/balance/internal/common/handler"
)

type IndexHandler struct {
	*handler.Handler

	itemService item.Service
	component   *component.Component
}

func NewIndexHandler(
	itemService item.Service,
) *IndexHandler {
	return &IndexHandler{
		Handler:     handler.New(),
		itemService: itemService,
		component:   component.New(),
	}
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	month, year := r.URL.Query().Get("month"), r.URL.Query().Get("year")

	request := item.ListRequest{Month: month, Year: year}

	items, err := h.itemService.List(r.Context(), request)
	if err != nil {
		h.SetError(w, err)

		return
	}

	w.Header().Add(
		"Hx-Trigger-After-Swap",
		fmt.Sprintf(`{"balance.items.shown":{"month":"%s","year":"%s"}}`, month, year),
	)

	err = h.component.Index(r.URL.Query(), items).Render(w)

	h.SetError(w, err)
}
