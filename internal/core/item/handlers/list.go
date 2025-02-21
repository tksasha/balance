package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/core/common"
	indexpagecomponents "github.com/tksasha/balance/internal/core/indexpage/components"
	"github.com/tksasha/balance/internal/core/item"
	"github.com/tksasha/balance/internal/core/item/components"
)

type ListHandler struct {
	*common.BaseHandler

	itemService     item.Service
	itemsComponent  *components.ItemsComponent
	monthsComponent *indexpagecomponents.MonthsComponent
}

func NewListHandler(
	baseHandler *common.BaseHandler,
	itemService item.Service,
	itemsComponent *components.ItemsComponent,
	monthsComponent *indexpagecomponents.MonthsComponent,
) *ListHandler {
	return &ListHandler{
		BaseHandler:     baseHandler,
		itemService:     itemService,
		itemsComponent:  itemsComponent,
		monthsComponent: monthsComponent,
	}
}

func (h *ListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	items, err := h.handle(r)
	if err != nil {
		h.SetError(w, err)

		return
	}

	err = h.itemsComponent.Index(items, h.monthsComponent.Months(r.URL.Query())).Render(w)

	h.SetError(w, err)
}

func (h *ListHandler) handle(r *http.Request) (item.Items, error) {
	request := item.IndexRequest{
		Year:  r.URL.Query().Get("year"),
		Month: r.URL.Query().Get("month"),
	}

	return h.itemService.Index(r.Context(), request)
}
