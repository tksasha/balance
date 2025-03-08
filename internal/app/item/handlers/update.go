package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/tksasha/balance/internal/app/item"
	"github.com/tksasha/balance/internal/app/item/component"
	"github.com/tksasha/balance/internal/common"
	"github.com/tksasha/balance/internal/common/currency"
	"github.com/tksasha/balance/internal/common/handler"
	"github.com/tksasha/validation"
)

type UpdateHandler struct {
	*handler.Handler

	itemService     item.Service
	categoryService item.CategoryService
	component       *component.Component
}

func NewUpdateHandler(
	itemService item.Service,
	categoryService item.CategoryService,
) *UpdateHandler {
	return &UpdateHandler{
		Handler:         handler.New(),
		itemService:     itemService,
		categoryService: categoryService,
		component:       component.New(),
	}
}

func (h *UpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		h.SetError(w, common.ErrParsingForm)

		return
	}

	request := item.UpdateRequest{
		ID:          r.PathValue("id"),
		Date:        r.FormValue("date"),
		Formula:     r.FormValue("formula"),
		CategoryID:  r.FormValue("category_id"),
		Description: r.FormValue("description"),
	}

	item, err := h.itemService.Update(r.Context(), request)
	if err == nil {
		w.Header().Add(
			"Hx-Trigger-After-Swap",
			fmt.Sprintf(
				`{"balance.item.updated":{"currency":"%s","month":"%s","year":"%s"}}`,
				currency.GetCode(item.Currency),
				strconv.Itoa(int(item.Date.Month())),
				strconv.Itoa(item.Date.Year()),
			),
		)

		err := h.component.Update(item).Render(w)

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

		err = h.component.Edit(item, categories, verrors).Render(w)

		h.SetError(w, err)

		return
	}

	h.SetError(w, err)
}
