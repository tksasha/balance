package handlers

import (
	"net/http"

	indexcomponent "github.com/tksasha/balance/internal/app/index/component"
	"github.com/tksasha/balance/internal/app/item"
	"github.com/tksasha/balance/internal/app/item/components"
	"github.com/tksasha/balance/internal/common/handler"
)

type ListHandler struct {
	*handler.Handler

	itemService    item.Service
	itemsComponent *components.ItemsComponent
	indexComponent *indexcomponent.IndexComponent
}

func NewListHandler(
	itemService item.Service,
	itemsComponent *components.ItemsComponent,
	indexComponent *indexcomponent.IndexComponent,
) *ListHandler {
	return &ListHandler{
		Handler:        handler.New(),
		itemService:    itemService,
		itemsComponent: itemsComponent,
		indexComponent: indexComponent,
	}
}

func (h *ListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	items, err := h.handle(r)
	if err != nil {
		h.SetError(w, err)

		return
	}

	months := h.indexComponent.Months(r.URL.Query())

	years := h.indexComponent.Years(r.URL.Query())

	err = h.itemsComponent.List(items, months, years).Render(w)

	h.SetError(w, err)
}

func (h *ListHandler) handle(r *http.Request) (item.Items, error) {
	request := item.ListRequest{
		Year:  r.URL.Query().Get("year"),
		Month: r.URL.Query().Get("month"),
	}

	return h.itemService.List(r.Context(), request)
}
