package handlers

import (
	"net/http"

	indexcomponents "github.com/tksasha/balance/internal/app/index/components"
	"github.com/tksasha/balance/internal/app/item"
	"github.com/tksasha/balance/internal/app/item/components"
	"github.com/tksasha/balance/internal/common"
)

type ListHandler struct {
	*common.BaseHandler

	itemService     item.Service
	itemsComponent  *components.ItemsComponent
	monthsComponent *indexcomponents.MonthsComponent
	yearsComponent  *indexcomponents.YearsComponent
}

func NewListHandler(
	baseHandler *common.BaseHandler,
	itemService item.Service,
	itemsComponent *components.ItemsComponent,
	monthsComponent *indexcomponents.MonthsComponent,
	yearsComponent *indexcomponents.YearsComponent,
) *ListHandler {
	return &ListHandler{
		BaseHandler:     baseHandler,
		itemService:     itemService,
		itemsComponent:  itemsComponent,
		monthsComponent: monthsComponent,
		yearsComponent:  yearsComponent,
	}
}

func (h *ListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	items, err := h.handle(r)
	if err != nil {
		h.SetError(w, err)

		return
	}

	months := h.monthsComponent.Months(r.URL.Query())

	years := h.yearsComponent.Years(r.URL.Query())

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
