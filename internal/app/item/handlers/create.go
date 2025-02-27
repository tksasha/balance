package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/tksasha/balance/internal/app/item"
	"github.com/tksasha/balance/internal/app/item/components"
	"github.com/tksasha/balance/internal/common"
	"github.com/tksasha/balance/internal/common/handler"
	"github.com/tksasha/validation"
)

type CreateHandler struct {
	*handler.Handler

	itemService     item.Service
	categoryService item.CategoryService
	itemsComponent  *components.ItemsComponent
}

func NewCreateHandler(
	itemService item.Service,
	categoryService item.CategoryService,
	itemsComponent *components.ItemsComponent,
) *CreateHandler {
	return &CreateHandler{
		Handler:         handler.New(),
		itemService:     itemService,
		categoryService: categoryService,
		itemsComponent:  itemsComponent,
	}
}

func (h *CreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	resource, err := h.handle(r)
	if err == nil {
		listRequest := item.ListRequest{
			Month: strconv.Itoa(int(resource.Date.Month())),
			Year:  strconv.Itoa(resource.Date.Year()),
		}

		items, err := h.itemService.List(r.Context(), listRequest)
		if err != nil {
			h.SetError(w, err)

			return
		}

		err = h.itemsComponent.List(items, nil, nil).Render(w)

		h.SetError(w, err)

		return
	}

	var verrors validation.Errors
	if errors.As(err, &verrors) {
		categories, err := h.categoryService.List(r.Context())
		if err != nil {
			h.SetError(w, err)

			return
		}

		err = h.itemsComponent.Edit(resource, categories, verrors).Render(w)

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
