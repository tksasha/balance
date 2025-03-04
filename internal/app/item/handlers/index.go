package handlers

import (
	"net/http"

	indexcomponent "github.com/tksasha/balance/internal/app/index/component"
	"github.com/tksasha/balance/internal/app/item"
	"github.com/tksasha/balance/internal/app/item/component"
	"github.com/tksasha/balance/internal/common/handler"
)

type IndexHandler struct {
	*handler.Handler

	itemService    item.Service
	component      *component.Component
	indexComponent *indexcomponent.Component
}

func NewIndexHandler(
	itemService item.Service,
) *IndexHandler {
	return &IndexHandler{
		Handler:        handler.New(),
		itemService:    itemService,
		component:      component.New(),
		indexComponent: indexcomponent.New(),
	}
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	items, err := h.handle(r)
	if err != nil {
		h.SetError(w, err)

		return
	}

	months := h.indexComponent.Months(r.URL.Query())

	years := h.indexComponent.Years(r.URL.Query())

	err = h.component.List(items, months, years).Render(w)

	h.SetError(w, err)
}

func (h *IndexHandler) handle(r *http.Request) (item.Items, error) {
	request := item.ListRequest{
		Year:  r.URL.Query().Get("year"),
		Month: r.URL.Query().Get("month"),
	}

	return h.itemService.List(r.Context(), request)
}
