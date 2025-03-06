package handlers

import (
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
	items, err := h.handle(r)
	if err != nil {
		h.SetError(w, err)

		return
	}

	err = h.component.Index(items).Render(w)

	h.SetError(w, err)
}

func (h *IndexHandler) handle(r *http.Request) (item.Items, error) {
	request := item.ListRequest{
		Year:  r.URL.Query().Get("year"),
		Month: r.URL.Query().Get("month"),
	}

	return h.itemService.List(r.Context(), request)
}
