package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/app/item"
	"github.com/tksasha/balance/internal/app/item/components"
	"github.com/tksasha/balance/internal/common"
)

type EditHandler struct {
	*common.BaseHandler

	itemService     item.Service
	categoryService item.CategoryService
	itemsComponent  *components.ItemsComponent
}

func NewEditHandler(
	baseHandler *common.BaseHandler,
	itemService item.Service,
	categoryService item.CategoryService,
	itemsComponent *components.ItemsComponent,
) *EditHandler {
	return &EditHandler{
		BaseHandler:     baseHandler,
		itemService:     itemService,
		categoryService: categoryService,
		itemsComponent:  itemsComponent,
	}
}

func (h *EditHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	item, err := h.handle(w, r)
	if err != nil {
		h.SetError(w, err)

		return
	}

	categories, err := h.categoryService.List(r.Context())
	if err != nil {
		h.SetError(w, err)

		return
	}

	err = h.itemsComponent.Edit(item, categories).Render(w)

	h.SetError(w, err)
}

func (h *EditHandler) handle(_ http.ResponseWriter, r *http.Request) (*item.Item, error) {
	return h.itemService.Edit(r.Context(), r.PathValue("id"))
}
