package handlers

import (
	"errors"
	"net/http"

	"github.com/tksasha/balance/internal/app/category"
	"github.com/tksasha/balance/internal/app/item"
	"github.com/tksasha/balance/internal/app/item/components"
	"github.com/tksasha/balance/internal/common"
	"github.com/tksasha/balance/pkg/validation"
)

type CreateHandler struct {
	*common.BaseHandler

	itemService     item.Service
	categoryService category.Service
	itemsComponent  *components.ItemsComponent
}

func NewCreateHandler(
	baseHandler *common.BaseHandler,
	itemService item.Service,
	categoryService category.Service,
	itemsComponent *components.ItemsComponent,
) *CreateHandler {
	return &CreateHandler{
		BaseHandler:     baseHandler,
		itemService:     itemService,
		categoryService: categoryService,
		itemsComponent:  itemsComponent,
	}
}

func (h *CreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	item, err := h.handle(r)
	if err == nil {
		w.WriteHeader(http.StatusNoContent)

		return
	}

	var verrors validation.Errors
	if errors.As(err, &verrors) {
		categories, err := h.categoryService.List(r.Context())
		if err != nil {
			h.SetError(w, err)

			return
		}

		err = h.itemsComponent.Create(item, categories, verrors).Render(w)

		h.SetError(w, err)

		return
	}

	h.SetError(w, err)
}

func (h *CreateHandler) handle(r *http.Request) (*item.Item, error) {
	if err := r.ParseForm(); err != nil {
		return nil, common.ErrParsingForm
	}

	request := item.CreateRequest{
		Date:        r.FormValue("date"),
		Formula:     r.FormValue("formula"),
		CategoryID:  r.FormValue("category_id"),
		Description: r.FormValue("description"),
	}

	return h.itemService.Create(r.Context(), request)
}
